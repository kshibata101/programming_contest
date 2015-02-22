#include <iostream>
#include <vector>
using namespace std;
     
int main() {
int n, m, tmp;
 vector<int> p;
     
 cin >> n >> m;
 for (int i = 0; i < n; i++) {
   cin >> tmp;
   p.push_back(tmp);
 }
     
 int ans = 0;
 for (int i = 0; i < m; i++) {
   cin >> tmp;
   ans += p[tmp-1];
 }
     
 cout << ans << endl;
 return 0;
}

