#include<iostream>
#include<vector>
#include<algorithm>
using namespace std;

void print(int** d, int m, int n) {
    for(int i = 0; i < m; i++){
        for(int j = 0; j < n; j++){
            cout << d[i][j] << " ";
        }
        cout << endl;
    }
    cout << endl;
}

class Solution {
public:
    int minDistance(string word1, string word2) {
        int **dis = new int*[word1.size()+1];
        for(int i = 0; i < word1.size()+1; i++){
            dis[i] = new int[word2.size()+1];
        }

        for(int i = 0; i < word1.size()+1; i ++) {
            dis[i][0] = i;
        }
        for(int i = 0; i < word2.size()+1; i ++){
            dis[0][i] = i;
        }

        for(int i = 1; i < word1.size()+1; i ++){
            for(int j = 1; j < word2.size()+1; j++){
                char c1 = word1[i-1], c2 = word2[j-1];
                int d1 = dis[i-1][j] + 1;
                int d2 = dis[i][j-1] + 1;
                int d3 = c1 == c2 ? dis[i-1][j-1] : dis[i-1][j-1]+1;
                dis[i][j] = min(d1, min(d2, d3));
            }
        }

        return dis[word1.size()][word2.size()];
    }
};

int main(){
    Solution s;
    cout << s.minDistance("", "") << endl;
    cout << s.minDistance("horse", "ros") << endl;
    cout << s.minDistance("intention", "execution") << endl;
    return 0;
}
