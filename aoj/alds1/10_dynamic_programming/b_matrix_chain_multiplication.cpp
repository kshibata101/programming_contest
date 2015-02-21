#include <iostream>
#include <vector>

using namespace std;

struct matrix {
  int r;
  int c;
};

// 連鎖行列積問題
int main() {
  vector<matrix> mat;
  int n;
  cin >> n;
  for (int i = 0; i < n; i++) {
    int tmp1,tmp2;
    cin >> tmp1 >> tmp2;
    matrix m = {tmp1, tmp2};
    mat.push_back(m);
  }
  
  int** memo = new int*[n];
  for (int s = 0; s < n; s++) {
    memo[s] = new int[n];
  }

  for (int l = 1; l < n; l++) {
    for (int s = 0; s < n - l; s++) {
      
      int minval = (1 << 21);
      int sum = 0;
      for (int i = 0; i < l; i++) {
        sum = memo[s][s+i] + memo[s+i+1][s+l];
        matrix m1 = mat[s];
        matrix m2 = mat[s+i];
        matrix m3 = mat[s+l];

        sum += m1.r * m2.c * m3.c;

        minval = min(minval, sum);
      }
      memo[s][s+l] = minval;
    }
  }

  cout << memo[0][n-1] << endl;
  return 0;
}
