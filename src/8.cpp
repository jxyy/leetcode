#include<iostream>
using namespace std;

class Solution {
public:
    int myAtoi(string str) {
        if(str.length() == 0){
            return 0;
        }
        int n = 0;
        int index = 0;
        for(; index < str.length(); index++){
            if(str[index] != ' '){
                break;
            }
        }

        bool negetive = str[index]=='-'? true:false;
        for(int i = str[index]=='+'||str[index]=='-'?index+1:index; i < str.length(); i++){
            char c = int(str[i]) - 48;
            if(c>=0 && c<=9){
                if(!negetive){
                    if(n<214748364 || (n==214748364&&c<=7)){
                        n *= 10;
                        n += c;
                    }else{
                        n = 2147483647;
                    }
                }else{
                    if(n > 0){
                        n = -n;
                    }
                    if(n>-214748364 || (n==-214748364&&c<=8)){
                        n *= 10;
                        n -= c;
                    }else{
                        n = -2147483648;
                    }
                }
            }else{
                break;
            }
        }
        return n;
    }
};

int main(){
    Solution s;
    cout << s.myAtoi("    010") << endl;
    return 0;
}
