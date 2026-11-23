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

func Constructor() Twitter {
	return Twitter{follows: map[int]map[int]bool{}}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.tweets = append(this.tweets, tweet{userId, tweetId})
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	feed := []int{}

	// newest tweets are at the end of the list
	for i := len(this.tweets) - 1; i >= 0 && len(feed) < 10; i-- {
		author := this.tweets[i].userId
		if author == userId || this.follows[userId][author] {
			feed = append(feed, this.tweets[i].tweetId)
		}
	}
	return feed
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.follows[followerId] == nil {
		this.follows[followerId] = map[int]bool{}
	}
	this.follows[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.follows[followerId], followeeId)
}
