function twoSum(nums, target) {

    let sorted = nums
    let map = {}
    nums.map((a, i) => {
        map[a] = i
    })
    for (let i = 0; i < nums.length; i++) {
        let a = nums[i]
        let key = target - a
        if (map[key] && map[key] !== i) {
            return [i, map[key]].sort()
        }
    }
    return null


};

function twoSum2(nums, target) {
    for (let i = 0; i < nums.length; i++) {
        let j = nums.indexOf(target - nums[i]);
        if (i != j && j > -1) {
            return [i, j];
        }
    }
    return [];
};

const s = [3, 2, 4];
const t = 6;

const start = Date.now()
const pos = twoSum(s, t)
console.log('excursion', Date.now() - start, 'ms')
console.log('Answer', pos)