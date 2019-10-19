# About

# Section3
scoped styling,
dynamic styling,
sudo selector and media query in jsx with radium
* this is good way
sudo selector and media query with css loader in webpack.
### Radium
this lib allow to use css in jsx
To use media query entire react app need to be wrapped with 'StyleRoot'

### webpack config for css module
css loader will generate unique class name

* after `npm run eject`

* add
`importLoaders: 1,`
`modules: true,`
`localIdentName: '[name]__[local]__[hash:base65:5]'` at css loader option in webpack.config 

* import css as js object
`import classes from './*.css'`

* use jsx class
`<div className={classes.App}></div>`

* nested class
`.App button.Red {...}`
can be called as `classes.Red` in .js file

##### Global
`:global .Post {...}`

# Section 4
### Error boundary
higher order function. When there is error, can show something what you want

### Difference between stateful component and stateless component
| | stateful component | stateless component |
|state access | yes | no |
|life cycle access | yes | no |
|props | yes(via this.props) | yes(via props) |
|usage| Use only if need to manage state or access to lifecycle hooks| Use in all other cases|

