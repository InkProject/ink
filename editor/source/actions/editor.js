import { ACTION } from './index';
import yaml from 'js-yaml';
import _ from 'lodash';

let parseConfig = function(data, noContent) {
    let content;
    if (noContent) {
        content = data;
    } else {
        content = data.split('---')[0] || '\n';
    }
    try {
        return yaml.safeLoad(content);
    } catch (err) {
        console.log(err);
        return null;
    }
};

export function setHeader(data) {
    let config = parseConfig(data);
    return {
        type: ACTION.SET_HEADER,
        title: config.title,
        tags: config.tags
    }
}

export function setEditor(data) {
    let dataAry = data.split('---');
    let configData = dataAry[0];
    let content = _.trim(dataAry[1] || '');
    let config = parseConfig(configData);
    let { title, tags } = config || {};
    return {
        type: ACTION.SET_CONTENT,
        title,
        tags,
        config: configData,
        content
    }
}
