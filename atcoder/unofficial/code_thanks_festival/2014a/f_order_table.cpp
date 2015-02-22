#include <iostream>
#include <vector>
#define eol '\n';
using namespace std;

int upper_num(int id, vector<vector<int>>& up, vector<bool>& visit) {
  visit[id] = true;

  int r = 0;
  for (int i = 0; i < up[id].size(); i++) {
    if (!visit[up[id][i]]) {
      r += 1 + upper_num(up[id][i], up, visit);
    }
  }
  return r;
}

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);
  
  int n,m;
  cin >> n >> m;
  
  vector<vector<int>> up(n);
  
  for (int i = 0; i < m; i++) {
    int a,b;
    cin >> a >> b;
    a--;
    b--;
    
    up[b].push_back(a);
  }

  vector<bool> visit(n, false);

  cout << upper_num(0, up, visit) + 1 << eol;
  
  return 0;
}
