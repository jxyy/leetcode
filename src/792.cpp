#include<iostream>
#include<vector>
#include<unordered_map>
using namespace std;

void printv(vector<int>& v) {
    for(auto i : v) {
        cout << i << " ";
    }
    cout << endl;
}

class Solution {
public:
    int binary_search(vector<int>& v, int target, int start, int end) {
        // cout << target << " " << start << " " << end << endl;
        if (start >= end) {
            return end;
        }
        int m = start + (end - start) / 2;
        if (target == v[m]) {
            // no possiable here
            return m + 1;
        } else if (target < v[m]) {
            return binary_search(v, target, start, m);
        } else {
            return binary_search(v, target, m+1, end);
        }
    }

    int numMatchingSubseq(string S, vector<string>& words) {
        unordered_map<char, vector<int>> cmap;
        for(int i = 0; i < S.length(); i++) {
            cmap[S[i]].push_back(i);
        }
        
        int count = 0;
        for(auto w : words) {
            if (w.length() > S.length()) {
                continue;
            }
            int i = -1, j = 0;
            // cout << "word:" << w.length() << endl;
            for(; j < w.length(); j++) {
                vector<int> cindices = cmap[w[j]];
                int ci = binary_search(cindices, i, 0, cindices.size());
                // cout << "near i:" << ci << endl;
                if (ci < cindices.size() && cindices[ci] > i) {
                    i = cindices[ci];
                } else{
                    break;
                }
            }
            if (j == w.length()) {
                // cout << "666" << endl;
                count ++;
            }
        }        
        return count;
    }
};


int main(){
    Solution s;
    string S = "abcde";
    vector<string> words;
    words.push_back("a");
    words.push_back("bb");
    words.push_back("acd");
    words.push_back("ace");
    cout << s.numMatchingSubseq(S, words) << endl;

    words.clear();
    S = "rwpddkvbnnuglnagtvamxkqtwhqgwbqgfbvgkwyuqkdwhzudsxvjubjgloeofnpjqlkdsqvruvabjrikfwronbrdyyjnakstqjac";
    words.push_back("wpddkvbnn");
    words.push_back("lnagtva");
    words.push_back("wpddkvbnn");
    words.push_back("kvbnnuglnagtvamxkqtwhqgwbqgfbvgkwyuqkdwhzudsxvju");
    words.push_back("rwpddkvbnnugln");
    cout << s.numMatchingSubseq(S, words) << endl;
    return 0;
}
