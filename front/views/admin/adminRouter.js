import Dashboard from './dashboard'
import Admin from './admin'
import Files from './files'
import Library from './library'
import Users from './users'

export default [
    {
        path: '/admin',
        name: 'admin',
        component: Admin,
        title: 'Admin',
        redirect: { name: 'dashboard' },
        children: [
            {
                path: 'dashboard',
                name: 'dashboard',
                component: Dashboard
            },
            {
                path: 'files',
                name: 'files',
                component: Files
            },
            {
                path: 'library',
                name: 'library',
                component: Library
            },
            {
                path: 'users',
                name: 'users',
                component: Users
            }
        ]
    }
]