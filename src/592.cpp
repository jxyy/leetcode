#include<iostream>
#include<vector>
#include<string>
#include<math.h>
using namespace std;

class Solution {
public:
    int gcd(int a, int b) {
        // cout << "gcd" << a << " " << b << endl;
        int m, n;
        if (a > b) {
            m = a;
            n = b;
        } else {
            m = b;
            n = a;
        }
        return m % n == 0 ? n : gcd(n, m % n);
    }

    string fractionAddition(string expression) {
        int sum_a = 0, sum_b = 1;
        int next_a = 0, next_b = 0;
        int i = 0;
        int status = 0; // 0,1,2 parse_a, parse_b, calculate
        bool is_add = true;
        if (expression.length() > 0 && expression[0] == '-') {
            is_add = false;
            i ++;
        }
        for(; i < expression.length(); i ++) {
            char c = expression[i];
            if(status == 0) {
                if (c == '/'){
                    status = 1;
                } else {
                    next_a *= 10;
                    next_a += (c-48);
                }
            }else if(status == 1) {
                if (c == '+' || c == '-'){
                    if (is_add) {
                        sum_a = sum_a*next_b + sum_b*next_a;
                    } else {
                        sum_a = sum_a*next_b - sum_b*next_a;
                    }
                    sum_b = sum_b*next_b;
                    if( sum_a != 0) {
                        int cd = gcd(abs(sum_a), abs(sum_b));
                        sum_a /= cd;
                        sum_b /= cd;
                    } else {
                        sum_b = 1;
                    }
                    
                    status = 0;
                    next_a = 0;
                    next_b = 0;
                    is_add = (c == '+');
                } else {
                    next_b *= 10;
                    next_b += (c-48);
                }
            }
        }

        if (is_add) {
            sum_a = sum_a*next_b + sum_b*next_a;
        } else {
            sum_a = sum_a*next_b - sum_b*next_a;
        }
        sum_b = sum_b*next_b;
        if (sum_a != 0) {
            int c = gcd(abs(sum_a), abs(sum_b));
            sum_a /= c;
            sum_b /= c;
        } else {
            sum_b = 1;
        }
        return to_string(sum_a) + "/" + to_string(sum_b);
    }
};

int main(){
    Solution s;
    cout << s.fractionAddition("-1/2+1/2") << endl;
    cout << s.fractionAddition("-1/2+1/2+1/3") << endl;
    cout << s.fractionAddition("1/3-1/2") << endl;
    cout << s.fractionAddition("5/3+1/3") << endl;
    return 0;
}
