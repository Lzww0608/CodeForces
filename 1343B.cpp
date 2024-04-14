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
		if (n % 4 != 0) {
			cout << "NO" << endl;
			continue;
		}
		cout << "YES" << endl;
		for (int i = 1, x = 2; i <= n / 2; i++) {
			cout << x << " ";
			if ((i & 1) == 0) {
				x += 4;
			}
			else {
				x += 2;
			}
		}

		for (int i = 1, x = 1; i <= n / 2; i++) {
			cout << x << " ";
			if ((i & 1) == 0) {
				x += 2;
			}
			else {
				x += 4;
			}
		}
		cout << endl;
	}
	

	return 0;
}