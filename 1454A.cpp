#include <iostream>
#include <vector>
#include <unordered_map>
#include <cstdlib>
#include <cmath>
#include <algorithm>
#include <string>
using namespace std;
using ll = long long;


int main() {
	std::ios::sync_with_stdio(false);
	cin.tie(nullptr);
	cout.tie(nullptr);
	
	int t, n;
	cin >> t;
	while (t--) {
		cin >> n;
		if (n & 1) {
			for (int i = 1, t = n - 1; i < n; ++i) {
				cout << t << " ";
				if (i & 1) {
					t++;
				}
				else {
					t -= 3;
				}
			}
			cout << 1;
		}
		else {
			for (int i = n; i >= 1; i--) {
				cout << i << " ";
			}
		}
		cout << endl;
	}
	

	return 0;
}