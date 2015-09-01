#include<iostream>
#include<stack>
using namespace std;

class Solution {
public:

    bool isMatch(string s, string p){
        s += '~';
        p += '~';
        int star_count = 0;
        for(int i = 0; i < p.length(); i++){
            if(p[i] == '*'){
                star_count ++;
            }
        }
        int repeat_total = s.length() - (p.length() - 2*star_count);
        if(repeat_total < 0){
            return false;
        }
        stack<int> ss, sp, sr;
        ss.push(0);
        sp.push(0);
        sr.push(0);
        int i = 0, j = 0;
        int repeat_count = 0;
        while(!ss.empty()){
            i = ss.top();
            j = sp.top();
            repeat_count = sr.top();
            ss.pop();
            sp.pop();
            sr.pop();
            while(j < p.length() && i < s.length()){
                char c = s[i], pc = p[j];
                bool is_repeat = j+1 < p.length() && p[j+1] == '*';
                if(is_repeat){
                    if(c != pc && pc != '.'){
                        j += 2;
                    }else{
                        if(repeat_count < repeat_total){
                            ss.push(i+1);
                            sp.push(j);
                            sr.push(repeat_count+1);
                        }
                        j += 2;
                    }
                }else{
                    if(c != pc && pc != '.'){
                        break;
                    }else{
                        j ++;
                        i ++;
                    }
                }
                if(i >= s.length()){
                    break;
                }
            }

            if((i == s.length() && j == p.length()) || (i == s.length() && j+2 == p.length() && p[j+1] == '*')){
                return true;
            }
        }
        return false;
    }

    bool isMatch3(string s, string p){
        s += '~';
        p += '~';
        int star_count = 0;
        for(int i = 0; i < p.length(); i++){
            if(p[i] == '*'){
                star_count ++;
            }
        }
        int repeat_total = s.length() - (p.length() - 2*star_count);
        if(repeat_total < 0){
            return false;
        }
        stack<int> ss, sp, sr;
        ss.push(0);
        sp.push(0);
        sr.push(0);
        int i = 0, j = 0;
        int repeat_count = 0;
        while(!ss.empty()){
            i = ss.top();
            j = sp.top();
            repeat_count = sr.top();
            ss.pop();
            sp.pop();
            sr.pop();
            while(j < p.length() && i < s.length()){
                char c = s[i], pc = p[j];
                bool is_repeat = j+1 < p.length() && p[j+1] == '*';
                if(is_repeat){
                    if(c != pc && pc != '.'){
                        j += 2;
                    }else{
                        if(repeat_count < repeat_total){
                            ss.push(i);
                            sp.push(j+2);
                            sr.push(repeat_count);
                            i ++;
                            repeat_count ++;
                        }else{
                            j += 2;
                        }
                    }
                }else{
                    if(c != pc && pc != '.'){
                        break;
                    }else{
                        j ++;
                        i ++;
                    }
                }
                if(i >= s.length()){
                    break;
                }
            }

            if((i == s.length() && j == p.length()) || (i == s.length() && j+2 == p.length() && p[j+1] == '*')){
                return true;
            }
        }
        return false;
    }
};

int main(){
    Solution s;
    cout << s.isMatch3("aa", "a*") << endl;
    cout << s.isMatch3("aa", ".*") << endl;
    cout << s.isMatch3("ab", ".*") << endl;
    cout << s.isMatch3("aab", "c*a*b") << endl;
    return 0;
}
