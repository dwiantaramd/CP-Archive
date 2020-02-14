#include <bits/stdc++.h>

using namespace std;

int main() {
    int value,i,temp;
    int x[10];
    i = 0;
    while(cin >> value) {
        x[i] = value;
        i++;
    }
    for(int idx = 0; idx <= i; idx++){
        temp = x[idx] % x[idx+1];
        while(temp != 0)
        {
            x[idx] = x[idx+1];
            x[idx+1] = temp;
            temp = x[idx] % x[idx+1];
        }
    }
}
