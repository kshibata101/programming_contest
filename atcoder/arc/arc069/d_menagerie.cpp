#include <iostream>
#include <vector>
#define eol '\n'
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int N;
  string s;

  cin >> N >> s;

  char sw[] = "SW";
  for (int i = 0; i < 4; i++) {
    char* str = new char[N];
    str[0] = sw[i/2];
    str[1] = sw[i%2];
    for (int j = 0; j < N; j++) {
      
    }
  }
  
  return 0;
}
