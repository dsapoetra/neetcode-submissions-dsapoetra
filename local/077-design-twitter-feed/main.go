// Problem:    Design Twitter
// Difficulty: Medium
// Link:       https://neetcode.io/problems/design-twitter-feed
// Source:     Data Structures & Algorithms/design-twitter-feed/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Build a small Twitter:
//
//	PostTweet(userId, tweetId)        the user posts a tweet
//	GetNewsFeed(userId)               the 10 newest tweet ids from the user
//	                                  and from the people they follow,
//	                                  newest first
//	Follow(followerId, followeeId)    start following
//	Unfollow(followerId, followeeId)  stop following
//
// Example:
//
//	Input:  PostTweet(1, 10), PostTweet(2, 20), GetNewsFeed(1), GetNewsFeed(2),
//	        Follow(1, 2), GetNewsFeed(1), GetNewsFeed(2),
//	        Unfollow(1, 2), GetNewsFeed(1)
//	Output: -, -, [10], [20], -, [20, 10], [20], -, [10]
//
// How to run (from the "local" folder):
//
//	go run ./077-design-twitter-feed
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: keep ONE list with all tweets in the order they were posted, and
// for every user the set of people they follow.
// To build a news feed, walk through the tweet list from the newest to
// the oldest and take the tweets written by the user or by someone they
// follow, until we have 10.
// PostTweet, Follow, Unfollow: O(1). GetNewsFeed: O(number of tweets).
// (Faster versions merge per-user lists with a heap, but need much more code.)

type tweet struct {
	userId  int
	tweetId int
}

type Twitter struct {
	tweets  []tweet              // all tweets, oldest first
	follows map[int]map[int]bool // follows[a][b] = true means a follows b
}

// Constructor makes an empty Twitter: no tweets, nobody follows anybody.
func Constructor() Twitter {
	return Twitter{follows: map[int]map[int]bool{}}
}

// PostTweet adds the tweet at the end of the list, so the list stays
// ordered from oldest to newest.
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.tweets = append(this.tweets, tweet{userId, tweetId})
}

// GetNewsFeed returns the (at most) 10 newest tweet ids from the user
// and from the people the user follows, newest first.
func (this *Twitter) GetNewsFeed(userId int) []int {
	feed := []int{}

	// newest tweets are at the end of the list
	for i := len(this.tweets) - 1; i >= 0 && len(feed) < 10; i-- {
		author := this.tweets[i].userId
		// a user always sees their own tweets, plus those of people they follow
		// (a missing map entry gives false, so unknown users are no problem)
		if author == userId || this.follows[userId][author] {
			feed = append(feed, this.tweets[i].tweetId)
		}
	}
	return feed
}

// Follow makes followerId see the tweets of followeeId.
func (this *Twitter) Follow(followerId int, followeeId int) {
	// the first time this user follows someone: create their set
	if this.follows[followerId] == nil {
		this.follows[followerId] = map[int]bool{}
	}
	this.follows[followerId][followeeId] = true
}

// Unfollow removes followeeId from the set. Nothing happens if followerId
// was not following them (delete on a missing entry is safe).
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.follows[followerId], followeeId)
}

// ===================== Test cases =====================

func main() {
	// Twitter remembers everything, so the lines starting with "--" show
	// every action in the output.
	twitter := Constructor()

	// Example from the problem.
	twitter.PostTweet(1, 10)
	twitter.PostTweet(2, 20)
	fmt.Println(`-- twitter.PostTweet(1, 10); twitter.PostTweet(2, 20)`)
	check("user 1 sees only their own tweet", `twitter.GetNewsFeed(1)`,
		twitter.GetNewsFeed(1), []int{10})
	check("user 2 sees only their own tweet", `twitter.GetNewsFeed(2)`,
		twitter.GetNewsFeed(2), []int{20})

	twitter.Follow(1, 2)
	fmt.Println(`-- twitter.Follow(1, 2)`)
	check("after following: newest first", `twitter.GetNewsFeed(1)`,
		twitter.GetNewsFeed(1), []int{20, 10})
	check("following is one-way", `twitter.GetNewsFeed(2)`,
		twitter.GetNewsFeed(2), []int{20})

	twitter.Unfollow(1, 2)
	fmt.Println(`-- twitter.Unfollow(1, 2)`)
	check("after unfollowing", `twitter.GetNewsFeed(1)`,
		twitter.GetNewsFeed(1), []int{10})

	// A user who never posted and follows nobody.
	check("user without tweets", `twitter.GetNewsFeed(99)`,
		twitter.GetNewsFeed(99), []int{})

	// Unfollowing someone we do not follow must not crash.
	twitter.Unfollow(5, 6)
	fmt.Println(`-- twitter.Unfollow(5, 6)   (user 5 never followed user 6)`)
	check("still an empty feed", `twitter.GetNewsFeed(5)`,
		twitter.GetNewsFeed(5), []int{})

	// The feed shows at most 10 tweets: the 10 newest ones.
	busy := Constructor()
	for id := 1; id <= 12; id++ {
		busy.PostTweet(7, id)
	}
	fmt.Println(`-- busy := Constructor(); user 7 posts the tweets 1, 2, ..., 12`)
	check("only the 10 newest tweets", `busy.GetNewsFeed(7)`,
		busy.GetNewsFeed(7), []int{12, 11, 10, 9, 8, 7, 6, 5, 4, 3})

	// Tweets of two users are mixed by time.
	busy.PostTweet(8, 100)
	busy.PostTweet(7, 13)
	busy.Follow(8, 7)
	fmt.Println(`-- busy.PostTweet(8, 100); busy.PostTweet(7, 13); busy.Follow(8, 7)`)
	check("mixed by time", `busy.GetNewsFeed(8)`,
		busy.GetNewsFeed(8), []int{13, 100, 12, 11, 10, 9, 8, 7, 6, 5})

	report()
}

// ===================== Helpers for the tests =====================

// passed and failed count the test results.
var passed, failed int

// check runs one test case. It prints the input, the result we got and the
// result we expected, with PASS or FAIL in front.
//
//	name:  a short description of the case
//	input: the call we are testing, written as text so it can be printed
//	got:   the answer the solution returned
//	want:  the answer we expect
func check(name string, input string, got interface{}, want interface{}) {
	status := "PASS"
	if fmt.Sprint(got) == fmt.Sprint(want) {
		passed++
	} else {
		failed++
		status = "FAIL"
	}
	fmt.Printf("%s  %s\n", status, name)
	fmt.Printf("      input:    %s\n", input)
	fmt.Printf("      result:   %s\n", show(got))
	fmt.Printf("      expected: %s\n\n", show(want))
}

// show turns a value into text. Strings get quotes, so an empty string is visible.
func show(value interface{}) string {
	if text, ok := value.(string); ok {
		return fmt.Sprintf("%q", text)
	}
	return fmt.Sprint(value)
}

// report prints the total. The program exits with an error if a test failed.
func report() {
	fmt.Printf("%d passed, %d failed\n", passed, failed)
	if failed > 0 {
		os.Exit(1)
	}
}
