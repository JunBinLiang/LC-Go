using namespace std;
using ll = long long;
#define pb push_back
#define ve vector
#define FOR(i, a, b) for (int i = a; i < b; ++i)
#define RFOR(i, a, b) for (int i = a; i >= b; i--)
#define f first
#define se second
#define W while
#define um unordered_map
#define us unordered_set
#define be begin
#define en end

class Solution {
public:
    bool matchReplacement(string s, string sub, vector<vector<char>>& m) {
        um<char, us<char>> f;
        for(int i = 0; i < m.size(); i++) {
            char c1 = m[i][0];
            char c2 = m[i][1];
            f[c1].insert(c2);
        }
        
        
        for(int i = 0; i < s.size(); i++) {
            if(i + sub.size() - 1 >= s.size()) {
                break;
            }
            
            bool yes = true;
            for(int j = 0; j < sub.size(); j++) {
                if(sub[j] != s[i + j]) {
                    char c1 = sub[j], c2 = s[i + j];
                    if(f[c1].find(c2) == f[c1].end()) {
                        yes = false;
                        break;
                    }
                }
            }
            
            if(yes) return true;
        }
        
        
        return false;
    }
};