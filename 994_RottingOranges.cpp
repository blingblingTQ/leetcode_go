class Solution {
public:
    
        
    vector<pair<int, int> > freshNeighbors(vector<vector<int>>& grid, int row, int col) {
        int rows = grid.size();
        int cols = rows > 0 ? grid[0].size() : 0;
        vector<pair<int, int> > neighbors;
        if(row-1 >= 0 && grid[row-1][col] == 1) {
            neighbors.push_back(make_pair(row-1, col));
        }
        if(row+1 < rows && grid[row+1][col] == 1) {
            neighbors.push_back(make_pair(row+1, col));
        }
        if(col-1 >= 0 && grid[row][col-1] == 1) {
            neighbors.push_back(make_pair(row, col-1));
        }
        if(col+1 < cols && grid[row][col+1] == 1) {
            neighbors.push_back(make_pair(row, col+1));
        }
        return neighbors;
    }
    
    int orangesRotting(vector<vector<int>>& grid) {
        int rows = grid.size();
        int cols = rows > 0 ? grid[0].size() : 0;
        
        queue<pair<int, int> > q;
        int freshCount = 0;
        for(int i = 0; i < rows; i++) {
            for(int j = 0; j < cols; j++) {
                if(grid[i][j] == 2) {
                    q.push(make_pair(i, j));
                } else if(grid[i][j] == 1) {
                    freshCount++;
                }
            }
        }
        
        int minute = 0;
        while(!q.empty()) {
            if(freshCount == 0) {
                return minute;
            }  
            int size = q.size();
            minute++;
            for(int i = 0; i < size; i++) {
                auto location = q.front();
                q.pop();
                auto neighbors = freshNeighbors(grid, location.first, location.second);
                for(auto neighbor: neighbors) {
                    grid[neighbor.first][neighbor.second] = 2;
                    freshCount--;
                    q.push(neighbor);
                } 
            }
        }
        if(freshCount > 0) {
            return -1;
        }
        return minute;
    }
    
};
