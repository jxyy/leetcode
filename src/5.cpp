#include<iostream>
#include<string>
using namespace std;

class Solution {
public:
    string longestPalindrome(string s) {
        int max = 0, max_endinghere = 0;
        int start = 0;
        int ending_start = 0;
        bool allsame = true;
        for(int i = 0; i < s.length(); i++){
            char c = s[i];
            if(max_endinghere > 0){
                if(allsame){
                    if(c == s[i-1]){
                        max_endinghere++;
                    }else{
                        if(ending_start-1>=0 && c==s[ending_start-1]){
                            allsame = false;
                            ending_start--;
                            max_endinghere += 2;
                        }else{
                            i = ending_start + max_endinghere/2 + 1;
                            if(s[i-1] == s[i]){
                                ending_start = i-1;
                                max_endinghere = 2;
                            }else{
                                ending_start = i;
                                max_endinghere = 1;
                            }
                        }
                    }
                }else{
                    if(ending_start-1>=0 && c==s[ending_start-1]){
                        ending_start--;
                        max_endinghere += 2;
                    }else{
                        i = ending_start + max_endinghere/2 + 1;
                        if(s[i-1] == s[i]){
                            ending_start = i-1;
                            max_endinghere = 2;
                        }else{
                            ending_start = i;
                            max_endinghere = 1;
                        }
                        allsame = true;
                    }
                }
            }else{
                max_endinghere++;
            }
            
            if(max_endinghere >= max){
                max = max_endinghere;
                start = ending_start;
            }
            //cout << c << ':' << ending_start << ' ' << max_endinghere << ' ' << allsame << endl;
        }
        return s.substr(start, max);
    }
};

int main(){
    string s = "bananas";
    Solution sl;
    cout << sl.longestPalindrome(s) << endl;
    return 0;
}
