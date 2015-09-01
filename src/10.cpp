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
        int pi = 0;
        bool match = false;
        int tail_len_s = 0, tail_len_p = 0;
        for(int i = p.length()-1, j = s.length()-1; i >= 0 && j >= 0; i--, j--){
            if(p[i] == '*'){
                if(p[i-1] == s[j] || p[i-1] == '.'){
                    break;
                }else{
                    j++;
                    i--;
                    tail_len_p += 2;
                }
            }else if(p[i] == '.' || p[i] == s[j]){
                tail_len_s ++;
                tail_len_p ++;
            }else{
                return false;
            }
        }
        p = p.substr(0, p.length()-tail_len_p);
        s = s.substr(0, s.length()-tail_len_s);
        cout << "p:" << p << endl;
        cout << "s:" << s << endl;
        for(int i = 0; i < s.length(); i++){
            if(pi == p.length() && p[pi-1] != '*'){
                cout << "false 1" << endl;
                return false;
            }

            if(p[pi] == '.'){
                pi ++;
            }else if(p[pi] == '*'){
                if(p[pi-1] == s[i] || p[pi-1] == '.'){
                    
                }else if(pi+1 < p.length() && (p[pi+1] == s[i] || p[pi+1] == '.')){
                    pi += 2;
                }else{
                    pi = 0;
                    cout << "false 2" << endl;
                    return false;
                }
            }else if(p[pi] == s[i]){
                pi ++;
            }else if(pi+1 < p.length() && p[pi+1] == '*'){
                pi ++;
                i --;
            }else{
                pi = 0;
                cout << "false 3" << endl;
                return false;
            }
        }

        while( pi+1 < p.length()){
            if(pi >= p.length()-1){
                break;
            }else{
                if(p[pi+1] == '*'){
                    pi += 2;
                }else{
                    break;
                }
            }
        }

        if(pi >= p.length() || (pi == p.length()-1 && p[pi] == '*')){
            match = true;
        }
        cout << pi << endl;

        return match;
    }

};

int main(){
    Solution s;
    cout << s.isMatch("aa", "a*") << endl;
    cout << s.isMatch("aa", ".*") << endl;
    cout << s.isMatch("ab", ".*") << endl;
    cout << s.isMatch("aab", "c*a*b") << endl;
    return 0;
}
