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

	int n, t;
	cin >> n >> t;
	string s;
	cin >> s;
	bool f = true;
	for (int i = 0; i < t && f; i++) {
		f = false;
		for (int j = 0; j < s.length() - 1; j++) {
			if (s[j] == 'B' && s[j + 1] == 'G') {
				swap(s[j], s[j + 1]);
				f = true;
				j++;
			}
		}
	}
	
	cout << s << endl;
	return 0;
}