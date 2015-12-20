import React from 'react';
import classNames from 'classnames';

export default class extends React.Component {
    render() {
        const list = this.props.list;
        const onOpenArticle = this.props.onOpenArticle;
        return (
            <ul id="list">{
                list.get('data').map(item =>
                    <li className={
                        classNames('item hover', {
                            selected: item.get('id') == list.get('selected')})
                        } key={item.get('id')} onClick={onOpenArticle.bind(this, item.get('id'))}>
                        <div className="head">
                            <span className="date">2015-03-01 18:00</span>
                        </div>
                        <div className="title">{item.get('title')}</div>
                        <div className="preview">{item.get('preview')}</div>
                    </li>
                )
            }</ul>
        );
    }
}
