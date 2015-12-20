function showLoading(flag) {
    return {
        type: 'SHOW_LOADING',
        flag
    }
}

function refreshList(data) {
    return {
        type: 'REFRESH_LIST',
        data
    }
}

export function fetchList() {
    return dispatch => {
        dispatch(showLoading(true));
        fetch('http://localhost:8001/api/articles').then(function(response) {
            return response.json();
        }).then(function(data) {
            dispatch(refreshList(data));
            dispatch(showLoading(false));
        }).catch(function(error) {
            alert(JSON.stringify(error));
        });
    };
}

export function showList() {
    return {
        type: 'SHOW_LIST'
    }
}

export function hideList() {
    return {
        type: 'HIDE_LIST'
    }
}

function setContent(content) {
    return {
        type: 'SET_CONTENT',
        content
    }
}

function selectArticle(id) {
    return {
        type: 'SELECT_ARTICLE',
        id
    }
}

export function openArticle(id) {
    return dispatch => {
        dispatch(showLoading(true));
        fetch('http://localhost:8001/api/articles/' + id).then(function(response) {
            return response.json();
        }).then(function(data) {
            dispatch(setContent(data));
            dispatch(selectArticle(id));
            dispatch(showLoading(false));
        }).catch(function(error) {
            alert(JSON.stringify(error));
        });
    };
}
