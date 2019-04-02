class Node {
public:
    int val;
    vector<Node*> children;

    Node() {}

    Node(int _val, vector<Node*> _children) {
        val = _val;
        children = _children;
    }
};

class Solution {
public:

    //dfs
    int maxDepth(Node* root) {
        if (root == NULL) {
            return 0;
        }
        
        int max = 0;
        for(auto child : root->children) {
            max = std::max(max, maxDepth(child));
        }
        return max+1;
    }

    //bfs
    int maxDepth(Node* root) {
        int depth = 0;
        if(root == NULL) {
            return depth;
        }
        queue<Node *> q;
        q.push(root);
        while(!q.empty()) {
            depth++;
            int size = q.size();
            for(int i = 0; i < size; i++) {
                Node *node = q.front();
                q.pop();
                for(int j = 0; j < node->children.size(); j++) {
                    q.push(node->children[j]);
                }
            }
        }
        return depth;
    }
};
