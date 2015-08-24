#include<iostream>
using namespace std;

class Solution {
public:
    string convert(string s, int numRows) {
        string converted_s = s;
        int group_size = numRows + (numRows-2);
        group_size = group_size<1? 1:group_size;
        int index2 = 0;
        for(int i = 0; i < numRows; i++){
            int index1 = i;
            if(index1 < s.length()){
                converted_s[index2] = s[index1];
                index2 ++;
            }
            while(index1 < s.length()){
                int offset = group_size-i*2;
                if(offset > 0){
                    index1 += offset;
                    if(index1 < s.length()){
                        converted_s[index2] = s[index1];
                        index2 ++;
                    }else{
                        break;
                    }
                }
                offset = group_size-offset;
                if(offset > 0){
                    index1 += offset;
                    if(index1 < s.length()){
                        converted_s[index2] = s[index1];
                        index2 ++;
                    }else{
                        break;
                    }
                }
            }

        }
        return converted_s;
    }
};

int main(){
    Solution s;
    cout << s.convert("A", 1) << endl;
    return 0;
}
