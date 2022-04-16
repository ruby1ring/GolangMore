# 4.11-4.17

* [X] [Reids基础操作](./redis.md)

[135.分发糖果](https://leetcode-cn.com/problems/candy/)

[860.柃檬水找零](https://leetcode-cn.com/problems/lemonade-change/)

```
class Solution {
public:
    bool lemonadeChange(vector<int>& bills) {
        int five=0,ten=0,twenty=0;
        for (int bill:bills){
            if (bill==5){
                five++;
            }else if(bill==10){
                if (five==0){
                    return false;
                }
                five--;
                ten++;
            }else{
                if (five>0 && ten>0){
                    five--;
                    ten--;
                }else if(five>=3){
                    five-=3;
                }else{
                    return false;
                }
            }
        }
        return true;
    }
};
```

[406.根据身高重建队列](https://leetcode-cn.com/problems/queue-reconstruction-by-height/)

```
class Solution {
public:
    static bool cmp(const vector<int>& a,const vector<int>& b){
        if (a[0]==b[0])return a[1]<b[1];
        return a[0]>b[0];
    }
    vector<vector<int>> reconstructQueue(vector<vector<int>>& people) {
        sort(people.begin(),people.end(),cmp);
        vector<vector<int>> queue;
        for (int i=0;i<people.size();i++){
            int position=people[i][1];
            queue.insert(queue.begin()+position,people[i]);
        }
        return queue;
    }
};
```

[452.用最少数量的箭引爆气球](https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/)

```
class Solution {
public:
    static bool cmp(const vector<int>& a,const vector<int>& b){
        return a[0]<b[0];
    }
    int findMinArrowShots(vector<vector<int>>& points) {
        if(points.size()==0)return 0;
        sort(points.begin(),points.end(),cmp);
        int result=1;
        for (int i=1;i<points.size();i++){
            if(points[i][0]>points[i-1][1]){
                result++;
            }else{
                points[i][1]=min(points[i-1][1],points[i][1]);
            }
        }
        return result;

    }
};
```
