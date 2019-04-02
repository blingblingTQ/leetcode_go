struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};
 
class Solution {
public:
    
    //dfs
    void dfs(TreeNode* root, vector<vector<int> >& result, int level) {
        if(root == NULL) {
            return;
        }
        if(result.size() == level) {
            vector<int> v;
            result.push_back(v);
        }
        
        result[level].push_back(root->val);
        dfs(root->left, result, level+1);
        dfs(root->right, result, level+1);
    }
    
    vector<vector<int>> levelOrderBottom(TreeNode* root) {
        vector<vector<int> > result;
        dfs(root, result, 0);
        reverse(result.begin(), result.end());
        return result;
    }


    //bfs
    vector<vector<int>> levelOrderBottom(TreeNode* root) {
        vector<vector<int> > result;
        queue<TreeNode *> q;
        if (root == nullptr) {
            return result;
        }
        
        q.push(root);
        while(!q.empty()) {
            int levelSize  = q.size();
            vector<int> v(levelSize);
            for(int i = 0; i < levelSize; i++) {
                TreeNode *node = q.front();
                q.pop();
                
                v[i] = node->val;
                if(node->left) {
                    q.push(node->left);
                }
                if(node->right) {
                    q.push(node->right);
                }
            }
            result.push_back(v);
        }
        reverse(result.begin(), result.end());
        return result;
    }
};
