#include<iostream>
#include<vector>
using namespace std;

class Solution {
public:
    int valueAtI(vector<vector<int>>& board, int i) {
        int N = board.size();
        int rr = (i - 1) / N;
        int r = N - 1 - rr;
        int cc = (i - 1) % N;
        int c = cc;
        if (rr % 2 == 1) {
            c = N - 1 - c;
        }
        return board[r][c];
    }

    int snakesAndLadders(vector<vector<int>>& board) {
        
    }
};

int main(){
    Solution s;
    vector<vector<int>> board;
    int N = 4;
    for(int i = 0; i < N; i ++){
        vector<int> row;
        for(int j = 0; j < N; j++) {
            row.push_back(i*N+j);
        }
        board.push_back(row);
    }

    for(int i = 0; i < N*N; i++){
        cout << s.valueAtI(board, i+1) << endl;
    }

    return 0;
}
