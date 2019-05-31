#include<iostream>
#include<vector>
using namespace std;

int t = 0;

class Solution {
    bool search(vector<vector<char>>& board, string word, int ci, pair<int, int> current, vector<vector<bool>>& used) {
        if(ci == word.size()-1) return true;

        char c = word[ci];
        int i = current.first, j = current.second;
        auto dirs = {pair<int, int>(0, 1), pair<int, int>(0, -1), pair<int, int>(1, 0), pair<int, int>(-1, 0)};
        for(auto dir : dirs) {
            int ii = i + dir.first, jj = j + dir.second;
            if(ii >= 0 && ii < board.size() && jj >= 0 && jj < board[0].size() && !used[ii][jj]){
                used[ii][jj] = true;
                if(board[ii][jj] == word[ci+1] && search(board, word, ci+1, pair<int, int>(ii, jj), used)) return true;
                used[ii][jj] = false;
            }
        }
        return false;
    }
public:
    bool exist(vector<vector<char>>& board, string word) {
        if(word.size() == 0) return true;

        vector<pair<int, int>> start;
        vector<vector<bool>> used;
        for(int i = 0; i < board.size(); i++) {
            used.push_back(vector<bool>());
            for(int j = 0; j < board[i].size(); j++) {
                if(board[i][j] == word[0]) start.push_back(pair<int, int>(i, j));
                used[i].push_back(false);
            }
        }

        for(auto s : start) {
            used[s.first][s.second] = true;
            if(search(board, word, 0, s, used)) return true;
            used[s.first][s.second] = false;
        }
        return false;
    }
};

int main(){
    const int m = 3, n = 4;
    char raw[m][n] = {
        {'A', 'B', 'C', 'D'},
        {'S', 'F', 'C', 'S'},
        {'A', 'D', 'E', 'E'}
    };
    vector<vector<char>> board;
    for(int i = 0; i < m; i++) {
        board.push_back(vector<char>());
        for(int j = 0; j < n; j++) {
            board[i].push_back(raw[i][j]);
        }
    }

    Solution s;
    cout << s.exist(board, "ABCCED") << endl;
    cout << s.exist(board, "SEE") << endl;
    cout << s.exist(board, "ABCB") << endl;
    return 0;
}
