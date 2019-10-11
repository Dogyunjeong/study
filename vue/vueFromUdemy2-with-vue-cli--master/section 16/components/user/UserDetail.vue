<template>
    <div>
        <h3>Some User Details</h3>
        <p>User loaded has Id: {{ $route.params.id }}</p>
        <router-link
                 tag="button" 
                 :to="link"
                 class="btn btn-primary">Edit User</router-link>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                link: { 
                    name: 'userEdit', 
                    params: { 
                        id: this.$route.params.id 
                    }, 
                    query: { locale: 'en', q: 100 },
                    hash: '#data'
                }
            }
        },
        // before next() called, can't access in vue instance
        beforeRouteEnter(to, from, next) {
            if (true) {
                next()
            } else {
                next(false)
            }
            // next(vm => {
            //     //  way to access vue instance after has been loaded
            //     vm.link;
            // });
        }
    }
</script>