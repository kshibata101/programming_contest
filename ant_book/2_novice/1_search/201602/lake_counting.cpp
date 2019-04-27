#include <iostream>
#include <vector>
#define eol '\n'
using namespace std;

int n,m;
vector<string> lake;

bool dfs(int i, int j) {
  if (i < 0 || n <= i || j < 0 || m <= j) {
    return false;
  }
  if (lake[i][j] == '.') {
    return false;
  }
  lake[i][j] = '.';

  for (int dx = -1; dx <= 1; dx++) {
    for (int dy = -1; dy <= 1; dy++) {
      dfs(i + dx, j + dy);
    }
  }
  return true;
}

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  cin >> n >> m;
  lake.resize(n);
  for (int i = 0; i < n; i++) {
    cin >> lake[i];
  }

  int cnt = 0;
  for (int i = 0; i < n; i++) {
    for (int j = 0; j < m; j++) {
      if (dfs(i,j)) {
        cnt++;
      }
    }
  }
  cout << cnt << eol;
  
  return 0;
}
