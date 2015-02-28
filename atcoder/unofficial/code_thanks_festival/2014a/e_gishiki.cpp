#include <iostream>
#include <vector>
#define eol '\n';
using namespace std;

struct RC {
  int ra;
  int rb;
  int ca;
  int cb;
  RC(int a, int b, int c, int d): ra(a),rb(b),ca(c),cb(d){}
};

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int R,C,M,N;
  int ra,rb,ca,cb;
  cin >> R >> C >> M >> N;

  int** map = new int*[R+1];
  for (int i = 1; i <= R; i++) {
    map[i] = new int[C+1];
  }

  vector<RC> rc_list;

  for (int i = 0; i < N; i++) {
    cin >> ra >> rb >> ca >> cb;

    RC rc(ra, rb, ca, cb);
    rc_list.push_back(rc);

    for (int r = ra; r <= rb; r++) {
      for (int c = ca; c <= cb; c++) {
        map[r][c]++;
      }
    }
  }

  int sum = 0;
  for (int r = 1; r <= R; r++) {
    for (int c = 1; c <= C; c++) {
      if (map[r][c] % 4 == 0) sum++;
    }
  }

  for (int i = 0; i < N; i++) {
    RC rc = rc_list[i];
    int sum_i = sum;
    for (int r = rc.ra; r <= rc.rb; r++) {
      for (int c = rc.ca; c <= rc.cb; c++) {
        if (map[r][c] % 4 == 0) {
          sum_i--;
        } else if (map[r][c] % 4 == 1) {
          sum_i++;
        }
      }
    }
    if (sum_i == M) {
      // this is answer.
      cout << (i+1) << eol;
    }
  }

  return 0;
}
