#include <iostream>
using namespace std;

/**
 * Compute temporary array to maintain size of suffix which is same as prefix
 * Time/space complexity is O(size of pattern)
 */
void computeTemporaryArray(string pattern, int lps[])
{
    int index = 0;
    for (int i = 1; i < pattern.size(); ) {
        if (pattern[i] == pattern[index]) {
            lps[i] = index + 1;
            index++;
            i++;
        } else {
            if (index != 0) {
                index = lps[index-1];
            } else {
                lps[i] = 0;
                i++;
            }
        }
    }
}

/**
 * KMP algorithm of pattern matching.
 */
int KMP(string text, string pattern, int lps[])
{
    computeTemporaryArray(pattern, lps);

    for (int i = 0; i < pattern.size(); i++)
        cout << lps[i] << " ";
    cout << endl;

    int i = 0;
    int j = 0;
    while (i < text.size() && j < pattern.size()) {
        if (text[i] == pattern[j]) {
            i++;
            j++;
        } else {
            if (j != 0) {
                j = lps[j-1];
            } else {
                i++;
            }
        }
    }
    if (j == pattern.size())
        return i - j;
    return -1;
}

int main()
{
    // string s = "bacbababaabcbab", p = "abababca";
    // string s = "BBC ABCDAB ABCDABCDABDE", p = "ABCDABD";
    string s = "abcxabcdabxabcdabcdabca", p = "abcdabca";

    // 1.暴力法 O(m*n)
    // cout << NaiveStringSearch(s, p); // 15

    // 2.KMP算法 O(m+n)
    int lps[100] = { 0 };
    cout << KMP(s, p, lps) << endl; // 15

    return 0; 
}
