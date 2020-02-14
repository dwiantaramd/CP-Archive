#include <bits/stdc++.h>

using namespace std;

int main() {
    double x, y, z;

    cin >> x;
    cin >> y;
    cin >> z;

    bool c1, c2, c3;

    c1 = x+y > z;
    c2 = x+z > y;
    c3 = y+z > x;

    if(c1 && c2 && c3) {
        cout << "TRUE";
    } else {
        cout << "FALSE";
    }
}
