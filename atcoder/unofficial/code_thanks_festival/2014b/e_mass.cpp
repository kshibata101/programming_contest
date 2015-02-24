#include <iostream>
#include <vector>
#define eol '\n'
using namespace std;
     
int dx[] = {1, 0, -1 ,0};
int dy[] = {0, 1, 0, -1};
     
int** map;
     
bool bfs(int r, int c, int rg, int cg) {
  if (map[r][c] == 0) return false;
  if (r == rg && c == cg) return true;
  map[r][c] = 0;
  for (int i = 0; i < 4; i++) {
    if (bfs(r+dy[i], c+dx[i], rg, cg)) return true;
  }
  return false;
}
     
int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);
     
  int R,C,rs,cs,rg,cg,N,r,c,h,w;
  cin >> R >> C >> rs >> cs >> rg >> cg >> N;
  map = new int*[R+2];
  for (int i = 0; i < R+2; i++) {
    map[i] = new int[C+2];
  }
  for (int i = 0; i < N; i++) {
    cin >> r >> c >> h >> w;
    for (int y = 0; y < h; y++) {
      for (int x = 0; x < w; x++) {
        map[r + y][c + x] = 1;
      }
    }
  }
  if (bfs(rs,cs,rg,cg)) {
    cout << "YES" << eol;
  } else {
    cout << "NO" << eol;
  }
     
  return 0;
}
