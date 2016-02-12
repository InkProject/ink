import React from 'react'
import Component from '../index'
import classNames from 'classnames'
import { connect } from 'react-redux'
import _ from 'lodash'

class Header extends Component {
    render() {
        const title = this.props.title
        const tags = this.props.tags
        const path = this.props.routing.location.pathname
        return (
            <div id="header" className={classNames({edit: this.props.edit})}>
                <div className={classNames('title', 'hover', {readonly: path == '/edit/config' || path == '/edit/help'})} onClick={this.props.onClick}><i className="fa fa-cog"></i><span>{title}</span></div>
                {tags.length ? <div className="info">{
                    tags.map(item =>
                        item ? <span className="tag">{item}</span> : null
                    )
                }</div> : null}
            </div>
        )
    }
}

export default connect(function(state) {
    return {
        list: state.list,
        menu: state.menu,
        routing: state.routing
    }
})(Header)
