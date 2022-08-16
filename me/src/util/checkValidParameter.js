import axios from 'axios';

export const CheckStatus = {
    LOADING: 0,
    SUCCESS: 1,
    FAIL: 2,
}

export var CheckValidPostId = (postId) => {
    return new Promise(function (resolve, reject) {
        if(postId == null || isNaN(Number(postId)) || parseInt(postId) < 0) {
            reject(Error("post id isn't valid number"))
        }else {
            let convertPostId = parseInt(postId)
            axios.get("/v2/check-valid-postId?postId="+convertPostId).then(() => {
                resolve()
            }).catch( err => {
                reject(err)
            })
        }
    })
}