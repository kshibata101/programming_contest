#include <iostream>
#include <queue>
using namespace std;

int N, M;
string* map;

typedef pair<int, int> P;

int INF = 10000000;
int dx[] = {1, 0, -1, 0};
int dy[] = {0, 1, 0, -1};

int solve() {
  // init
  queue<P> que;
  int** d = new int*[N];
  int sx, sy;
  int nx, ny;
  P p;

  for (int i = 0; i < N; i++) {
    d[i] = new int[M];
    for (int j = 0; j < M; j++) {
      if (map[i][j] == 'S') {
        que.push(P(i, j));
      } else {
        d[i][j] = INF;
      }
    }
  }

  // bfs  
  while(que.size()) {
    p = que.front();
    que.pop();

    if (map[p.first][p.second] == 'G') {
      break;
    }
    
    // 近傍を探す
    for (int i = 0; i < 4; i++) {
      nx = p.first + dx[i];
      ny = p.second + dy[i];

      if (nx < 0 || ny < 0 || nx >= N || ny >= M) {
        continue;
      }

      // いけるのにまだ訪れていない場所
      if (map[nx][ny] != '#' && d[nx][ny] == INF) {
        que.push(P(nx, ny));
        d[nx][ny] = d[p.first][p.second] + 1;
      }
    }
  }

  return d[p.first][p.second];
}

int main() {
  cin >> N >> M;

  map = new string[N];
  for (int i = 0; i < N; i++) {
    cin >> map[i];
  }

  int res = solve();
  cout << res << endl;

  return 0;
}
