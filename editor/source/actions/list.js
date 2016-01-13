import { ACTION, editorAction, utilAction } from './index';

export function showLoading(flag) {
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
        dispatch(utilAction.showTip('load'));
        utilAction.apiRequest('GET', `articles`).then(function(data) {
            dispatch(refreshList(data));
            dispatch(utilAction.showTip('hide'));
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

export function openArticle(id) {
    return dispatch => {
        dispatch(utilAction.showTip('load'));
        utilAction.apiRequest('GET', `articles/${id}`).then(function(data) {
            dispatch(editorAction.setEditor(id, data));
            dispatch(utilAction.showTip('hide'));
        });
    };
}
