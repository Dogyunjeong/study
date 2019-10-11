export const lengthAware = {
    computed: {
        countAndAddText() {
            return this.text + '(' + this.text.length + ')'
        }
    },
}