class Solution {
    public boolean isPalindrome(String s) {
        String noSpaces = s.replaceAll("\\s+", ""); 
        noSpaces = noSpaces.replaceAll("[^a-zA-Z0-9]", "");
        noSpaces = noSpaces.toLowerCase();

        String reversed = new StringBuilder(noSpaces).reverse().toString();

        System.out.println(noSpaces);
         System.out.println(reversed);
        if (noSpaces.equals(reversed)) {
            return true;
        } 

        return false;
    }
}
