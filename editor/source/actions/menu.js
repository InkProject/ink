import { ACTION } from './index';
import * as editorAction from './editor';

function showLoading(flag) {
    return {
        type: ACTION.SHOW_LOADING,
        flag
    }
}

function refreshList(data) {
    return {
        type: ACTION.REFRESH_LIST,
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
        type: ACTION.SHOW_LIST
    }
}

export function hideList() {
    return {
        type: ACTION.HIDE_LIST
    }
}

function selectArticle(id) {
    return {
        type: ACTION.SELECT_ARTICLE,
        id
    }
}

export function openArticle(id) {
    return dispatch => {
        dispatch(showLoading(true));
        fetch('http://localhost:8001/api/articles/' + id).then(function(response) {
            return response.json();
        }).then(function(data) {
            dispatch(editorAction.setContent(data));
            dispatch(selectArticle(id));
            dispatch(showLoading(false));
        }).catch(function(error) {
            alert(JSON.stringify(error));
        });
    };
}
