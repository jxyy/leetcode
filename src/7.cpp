#include<iostream>
using namespace std;

class Solution {
public:
    int reverse(int x) {
        int rx = 0;
        for(int n = x; n != 0; n /= 10){
            int a = n%10;
            if((rx>=0&&rx<214748364) || (rx>=0&&rx==214748364&&a<=7) ||
                (rx<=0&&rx>-214748364) || (rx<=0&&rx==-214748364&&a>=-8)){
                rx *= 10;
                rx += a;
            }else{
                rx = 0;
                break;
            }
        }
        return rx;
    }
};

int main(){
    Solution s;
    cout << s.reverse(1463847412) << endl;
    return 0;
}
