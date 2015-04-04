#include <iostream>
#define eol '\n';
using namespace std;

int k;
int n = 8;
int* row;
int* col;
int* dif;
int* sum;

bool search(int cnt) {
  if (cnt == n+1) {
    // end
    return true;
  }

  for (int i = 0; i < n; i++) {
    for (int j = 0; j < n; j++) {
      if (!row[i] && !col[j] && !dif[i-j+n-1] && !sum[i+j]) {
        row[i] = cnt;
        col[j] = cnt;
        dif[i-j+7] = cnt;
        sum[i+j] = cnt;

        if (search(cnt+1)) {
          return true;
        }

        row[i] = 0;
        col[j] = 0;
        dif[i-j+n-1] = 0;
        sum[i+j] = 0;
      }
    }
  }
  return false;
}

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  row = new int[n];
  col = new int[n];
  dif = new int[n*2-1];
  sum = new int[n*2-1];

  cin >> k;
  int cnt = 1;
  for (int i = 0; i < k; i++) {
    int r,c;
    cin >> r >> c;
    row[r] = cnt;
    col[c] = cnt;
    dif[r-c+n-1] = cnt;
    sum[r+c] = cnt;
    cnt++;
  }

  search(cnt);

  for (int i = 0; i < n; i++) {
    for (int j = 0; j < n; j++) {
      if (col[j] == row[i]) {
        cout << 'Q';
      } else {
        cout << '.';
      }
    }
    cout << eol;
  }

  return 0;
}
