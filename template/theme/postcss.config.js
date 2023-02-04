module.exports = {
    plugins: [
        require('autoprefixer')({
            overrideBrowserslist: ['last 5 version', '>1%', 'ios 7']
        }),
        require('precsss'),
        require("postcss-import")
    ]
}