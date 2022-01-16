// vetur.config.js
/** @type {import('vls').VeturConfig} */

// Reference: https://vuejs.github.io/vetur/reference/#example

module.exports = {
    // **optional** default: `{}`
    // override vscode settings part
    // Notice: It only affects the settings used by Vetur.
    settings: {
      "vetur.useWorkspaceDependencies": true,
      "vetur.experimental.templateInterpolationService": true
    },

    // Tell Vetur where in the repo the Vue project(s) are located
    projects: [
      './firebase/frontend',
    ]
  }