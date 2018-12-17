#include<iostream>
#include<vector>
#include<unordered_map>
#include<string>
using namespace std;

struct Node {
    char c;
    int index;
    Node** children;
    Node(char c) : c(c), index(-1) {
        this->children = new Node*[26];
        for(int i = 0; i < 26; i ++){
            this->children[i] = NULL;
        }
    };
};

class Solution {
public:
    vector<string> findAllConcatenatedWordsInADict(vector<string>& words){
        vector<string> ret;
        unordered_map<int, vector<int>> mem;
        for(int i = 0; i < words.size(); i++){
            int hash = 0;
            for(char c : words[i]) hash = hash * 131 + c;
            mem[hash].push_back(i);
        }

        for(int i = 0; i < words.size(); i++) {
            vector<int> candidates;
            candidates.push_back(0);
            while(candidates.size() > 0){
                int start = candidates[candidates.size()-1];
                candidates.pop_back();
                string word = words[i].substr(start);
                int hash = 0;
                bool done = false;
                for(char c : word) {
                    hash = hash * 131 + c;
                    if(mem.find(hash) == mem.end()) continue;
                    for(int index : mem[hash]) {
                        if(i == index) continue;
                        int next_start = start + words[index].size();
                        if(next_start == words[i].size()) {
                            ret.push_back(words[i]);
                            done = true;
                            break;
                        }
                        candidates.push_back(next_start);
                    }
                }
                if(done) break;
            }
        }
        return ret;
    }

    vector<string> findAllConcatenatedWordsInADict2(vector<string>& words){
        // unsafe solution
        vector<string> ret;
        unordered_map<int, int> mem;
        for(int i = 0; i < words.size(); i++){
            int hash = 0;
            for(char c : words[i]) {
                hash = hash * 131 + c;
            }
            mem[hash] = i;
        }

        for(int i = 0; i < words.size(); i++) {
            vector<int> candidates;
            candidates.push_back(0);
            while(candidates.size() > 0){
                int start = candidates[candidates.size()-1];
                candidates.pop_back();
                string word = words[i].substr(start);
                int hash = 0;
                bool done = false;
                for(char c : word) {
                    hash = hash * 131 + c;
                    if(mem.find(hash) == mem.end()){
                        continue;
                    }
                    int index = mem[hash];
                    if(i == index){
                        continue;
                    }
                    int next_start = start + words[index].size();
                    if(next_start == words[i].size()) {
                        ret.push_back(words[i]);
                        done = true;
                        break;
                    }
                    candidates.push_back(next_start);
                }
                if(done) {
                    break;
                }
            }
            
        }
        return ret;
    }
};

void print(vector<string> vs) {
    for(auto s : vs) {
        cout << s << endl;
    }
}

int main(){
    Solution s;
    vector<string> input({"cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"});
    // vector<string> input({"cat","cats", "dog","hippopotamuses","rat","ratcatdogcat"});
    print(s.findAllConcatenatedWordsInADict(input));
    return 0;
}
