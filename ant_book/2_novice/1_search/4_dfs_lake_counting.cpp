#include <iostream>
using namespace std;

int N, M;
string *lake;

bool dfs(int x, int y) {
  if (x < 0 || y < 0 || x > M || y > N)
    return false;

  if (lake[y][x] != 'W') 
    return false;

  lake[y][x] = '.';
  
  for (int i = -1; i <= 1; i++) {
    for (int j = -1; j <= 1; j++) {
      dfs(x+j, y+i);
    }
  }
  return true;
}

int solve() {
  int count = 0;
  for (int i = 0; i < N; i++) {
    for (int j = 0; j < M; j++) {
      if (dfs(j, i)) 
        count++;
    }
  }
  return count;
}

int main() {
  cin >> N >> M;

  lake = new string[N];
  for (int i = 0; i < N; i++) {
    cin >> lake[i];
  }

  // solve
  int res = solve();
  cout << res << endl;
  return 0;
}

