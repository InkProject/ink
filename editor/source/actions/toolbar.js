import { ACTION, utilAction, listAction, apiURL } from './index';

export function saveContent(id, name, content) {
    return dispatch => {
        dispatch(listAction.showLoading(true));
        fetch(`${apiURL}/articles/${id}`, {
            method: 'PUT',
            body: {
                name,
                content
            }
        }).then(function(response) {
            return response.json();
        }).then(function(data) {
            dispatch(utilAction.showTip('保存成功'));
            dispatch(listAction.showLoading(false));
        }).catch(function(error) {
            alert(JSON.stringify(error));
        });
    };
}
