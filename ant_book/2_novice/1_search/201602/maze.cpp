#include <iostream>
#include <vector>
#include <queue>
#define eol '\n'
using namespace std;

typedef pair<int,int> P;

int n,m;
vector<string> maze;
queue<P> que;
vector<vector<int>> visit;
  
int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  cin >> n >> m;
  
  maze.resize(n);
  visit.resize(n,vector<int>(m, -1));
  
  int sx,sy,gx,gy;
  for (int i = 0; i < n; i++) {
    cin >> maze[i];
    for (int j = 0; j < maze[i].size(); j++) {
      if (maze[i][j] == 'S') {
        sx = i;
        sy = j;
      } else if (maze[i][j] == 'G') {
        gx = i;
        gy = j;
      }
    }
  }

  que.push(P(sx,sy));
  visit[sx][sy] = 0;
  vector<int> dx = {0,0,-1,1};
  vector<int> dy = {-1,1,0,0};

  while(que.size()) {
    P now = que.front();
    que.pop();

    if (now.first == gx && now.second == gy)  {
      cout << visit[gx][gy] << eol;
      break;
    }

    for (int i = 0; i < dx.size(); i++) {
      int x = now.first  + dx[i];
      int y = now.second + dy[i];

      if (x < 0 || x >= n || y < 0 || y >= m) continue;

      if (visit[x][y] >= 0) continue;

      if (maze[x][y] == '.' || maze[x][y] == 'G') {
        que.push(P(x,y));
        visit[x][y] = visit[now.first][now.second] + 1;
      }
    }
  }
  
  return 0;
}
