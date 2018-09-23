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
        int N = board.size();
        int target = N * N;
        vector<int> min_steps(N*N+1, -1);
        min_steps[1] = 0;
        vector<int> pending, next_pending;
        pending.push_back(1);
        int steps = 0;
        while(pending.size() > 0) {
            steps ++;
            for(int x : pending) {
                for(int i = 0; i < 6; i++) {
                    int next = x + i + 1;
                    if (next == target) {
                        return steps;
                    }
                    int board_value = valueAtI(board, next);
                    if(min_steps[next] == -1 || (board_value != -1 && min_steps[board_value] == -1)) {
                        min_steps[next] = steps;
                        if(board_value != -1) {
                            if (board_value == target) {
                                return steps;
                            }
                            min_steps[board_value] = steps;
                            next_pending.push_back(board_value);
                        } else {
                            next_pending.push_back(next);
                        }
                    }
                }
            }
            swap(pending, next_pending);
            next_pending.clear();
        }
        return -1;
    }
};

int main(){
    int raw_board[8][8] = {
        {-1, -1, -1, 46, 47, -1, -1, -1},
        {51, -1, -1, 63, -1, 31, 21, -1},
        {-1, -1, 26, -1, -1, 38, -1, -1},
        {-1, -1, 11, -1, 14, 23, 56, 57},
        {11, -1, -1, -1, 49, 36, -1, 48},
        {-1, -1, -1, 33, 56, -1, 57, 21},
        {-1, -1, -1, -1, -1, -1, 2, -1},
        {-1, -1, -1, 8, 3, -1, 6, 56}};

    Solution s;
    vector<vector<int>> board;
    int N = 8;
    for(int i = 0; i < N; i ++){
        vector<int> row;
        for(int j = 0; j < N; j++) {
            row.push_back(raw_board[i][j]);
        }
        board.push_back(row);
    }

    cout << s.snakesAndLadders(board) << endl;
    return 0;
}
