/* eslint-disable */
/* prettier-ignore */
// @ts-nocheck
// Generated by unplugin-vue-router. ‼️ DO NOT MODIFY THIS FILE ‼️
// It's recommended to commit this file.
// Make sure to add this file to your tsconfig.json file as an "includes" or "files" entry.

/// <reference types="unplugin-vue-router/client" />

import type {
  // type safe route locations
  RouteLocationTypedList,
  RouteLocationResolvedTypedList,
  RouteLocationNormalizedTypedList,
  RouteLocationNormalizedLoadedTypedList,
  RouteLocationAsString,
  RouteLocationAsRelativeTypedList,
  RouteLocationAsPathTypedList,

  // helper types
  // route definitions
  RouteRecordInfo,
  ParamValue,
  ParamValueOneOrMore,
  ParamValueZeroOrMore,
  ParamValueZeroOrOne,

  // vue-router extensions
  _RouterTyped,
  RouterLinkTyped,
  RouterLinkPropsTyped,
  NavigationGuard,
  UseLinkFnTyped,

  // data fetching
  _DataLoader,
  _DefineLoaderOptions,
} from 'unplugin-vue-router/types'

declare module 'vue-router/auto/routes' {
  export interface RouteNamedMap {
    '/[...path]': RouteRecordInfo<'/[...path]', '/:path(.*)', { path: ParamValue<true> }, { path: ParamValue<false> }>,
    '/admin/billing/': RouteRecordInfo<'/admin/billing/', '/admin/billing', Record<never, never>, Record<never, never>>,
    '/admin/domains': RouteRecordInfo<'/admin/domains', '/admin/domains', Record<never, never>, Record<never, never>>,
    '/admin/domains/[id]': RouteRecordInfo<'/admin/domains/[id]', '/admin/domains/:id', { id: ParamValue<true> }, { id: ParamValue<false> }>,
    '/admin/domains/new': RouteRecordInfo<'/admin/domains/new', '/admin/domains/new', Record<never, never>, Record<never, never>>,
    '/admin/general/': RouteRecordInfo<'/admin/general/', '/admin/general', Record<never, never>, Record<never, never>>,
    '/admin/groups': RouteRecordInfo<'/admin/groups', '/admin/groups', Record<never, never>, Record<never, never>>,
    '/admin/groups/[id]': RouteRecordInfo<'/admin/groups/[id]', '/admin/groups/:id', { id: ParamValue<true> }, { id: ParamValue<false> }>,
    '/admin/groups/new': RouteRecordInfo<'/admin/groups/new', '/admin/groups/new', Record<never, never>, Record<never, never>>,
    '/admin/users': RouteRecordInfo<'/admin/users', '/admin/users', Record<never, never>, Record<never, never>>,
    '/admin/users/[id]': RouteRecordInfo<'/admin/users/[id]', '/admin/users/:id', { id: ParamValue<true> }, { id: ParamValue<false> }>,
    '/admin/users/new': RouteRecordInfo<'/admin/users/new', '/admin/users/new', Record<never, never>, Record<never, never>>,
    '/chat/[tab]': RouteRecordInfo<'/chat/[tab]', '/chat/:tab', { tab: ParamValue<true> }, { tab: ParamValue<false> }>,
    '/chat/[tab]/[id]/': RouteRecordInfo<'/chat/[tab]/[id]/', '/chat/:tab/:id', { tab: ParamValue<true>, id: ParamValue<true> }, { tab: ParamValue<false>, id: ParamValue<false> }>,
    '/chat/[tab]/[id]/edit': RouteRecordInfo<'/chat/[tab]/[id]/edit', '/chat/:tab/:id/edit', { tab: ParamValue<true>, id: ParamValue<true> }, { tab: ParamValue<false>, id: ParamValue<false> }>,
    '/drive/files/[id]': RouteRecordInfo<'/drive/files/[id]', '/drive/files/:id', { id: ParamValue<true> }, { id: ParamValue<false> }>,
    '/drive/my-drive/': RouteRecordInfo<'/drive/my-drive/', '/drive/my-drive', Record<never, never>, Record<never, never>>,
    '/drive/my-drive/folders/[id]': RouteRecordInfo<'/drive/my-drive/folders/[id]', '/drive/my-drive/folders/:id', { id: ParamValue<true> }, { id: ParamValue<false> }>,
    '/drive/shared-drives/': RouteRecordInfo<'/drive/shared-drives/', '/drive/shared-drives', Record<never, never>, Record<never, never>>,
    '/drive/shared-drives/[id]/': RouteRecordInfo<'/drive/shared-drives/[id]/', '/drive/shared-drives/:id', { id: ParamValue<true> }, { id: ParamValue<false> }>,
    '/drive/shared-drives/[id]/folders/[folderId]': RouteRecordInfo<'/drive/shared-drives/[id]/folders/[folderId]', '/drive/shared-drives/:id/folders/:folderId', { id: ParamValue<true>, folderId: ParamValue<true> }, { id: ParamValue<false>, folderId: ParamValue<false> }>,
    '/drive/shared-with-me/': RouteRecordInfo<'/drive/shared-with-me/', '/drive/shared-with-me', Record<never, never>, Record<never, never>>,
    '/drive/trash/': RouteRecordInfo<'/drive/trash/', '/drive/trash', Record<never, never>, Record<never, never>>,
    '/forgot-password': RouteRecordInfo<'/forgot-password', '/forgot-password', Record<never, never>, Record<never, never>>,
    '/invite': RouteRecordInfo<'/invite', '/invite', Record<never, never>, Record<never, never>>,
    '/login': RouteRecordInfo<'/login', '/login', Record<never, never>, Record<never, never>>,
    '/m/record/entries/': RouteRecordInfo<'/m/record/entries/', '/m/record/entries', Record<never, never>, Record<never, never>>,
    '/m/record/entries/[id]': RouteRecordInfo<'/m/record/entries/[id]', '/m/record/entries/:id', { id: ParamValue<true> }, { id: ParamValue<false> }>,
    '/mail/': RouteRecordInfo<'/mail/', '/mail', Record<never, never>, Record<never, never>>,
    '/record/attendances/': RouteRecordInfo<'/record/attendances/', '/record/attendances', Record<never, never>, Record<never, never>>,
    '/record/competences/': RouteRecordInfo<'/record/competences/', '/record/competences', Record<never, never>, Record<never, never>>,
    '/record/competences/[id]': RouteRecordInfo<'/record/competences/[id]', '/record/competences/:id', { id: ParamValue<true> }, { id: ParamValue<false> }>,
    '/record/entries/': RouteRecordInfo<'/record/entries/', '/record/entries', Record<never, never>, Record<never, never>>,
    '/record/entries/[id]': RouteRecordInfo<'/record/entries/[id]', '/record/entries/:id', { id: ParamValue<true> }, { id: ParamValue<false> }>,
    '/record/projects/': RouteRecordInfo<'/record/projects/', '/record/projects', Record<never, never>, Record<never, never>>,
    '/record/projects/[id]': RouteRecordInfo<'/record/projects/[id]', '/record/projects/:id', { id: ParamValue<true> }, { id: ParamValue<false> }>,
    '/record/projects/export': RouteRecordInfo<'/record/projects/export', '/record/projects/export', Record<never, never>, Record<never, never>>,
    '/record/projects/new': RouteRecordInfo<'/record/projects/new', '/record/projects/new', Record<never, never>, Record<never, never>>,
    '/record/reports/': RouteRecordInfo<'/record/reports/', '/record/reports', Record<never, never>, Record<never, never>>,
    '/record/reports/new': RouteRecordInfo<'/record/reports/new', '/record/reports/new', Record<never, never>, Record<never, never>>,
    '/record/students/': RouteRecordInfo<'/record/students/', '/record/students', Record<never, never>, Record<never, never>>,
    '/record/students/[id]': RouteRecordInfo<'/record/students/[id]', '/record/students/:id', { id: ParamValue<true> }, { id: ParamValue<false> }>,
    '/record/students/[id]/competences/': RouteRecordInfo<'/record/students/[id]/competences/', '/record/students/:id/competences', { id: ParamValue<true> }, { id: ParamValue<false> }>,
    '/record/students/[id]/competences/[cid]': RouteRecordInfo<'/record/students/[id]/competences/[cid]', '/record/students/:id/competences/:cid', { id: ParamValue<true>, cid: ParamValue<true> }, { id: ParamValue<false>, cid: ParamValue<false> }>,
    '/record/tags/': RouteRecordInfo<'/record/tags/', '/record/tags', Record<never, never>, Record<never, never>>,
    '/reset-password': RouteRecordInfo<'/reset-password', '/reset-password', Record<never, never>, Record<never, never>>,
    '/school/certificates': RouteRecordInfo<'/school/certificates', '/school/certificates', Record<never, never>, Record<never, never>>,
    '/school/certificates/[id]': RouteRecordInfo<'/school/certificates/[id]', '/school/certificates/:id', { id: ParamValue<true> }, { id: ParamValue<false> }>,
    '/school/certificates/new': RouteRecordInfo<'/school/certificates/new', '/school/certificates/new', Record<never, never>, Record<never, never>>,
    '/school/grades': RouteRecordInfo<'/school/grades', '/school/grades', Record<never, never>, Record<never, never>>,
    '/school/grades/[id]': RouteRecordInfo<'/school/grades/[id]', '/school/grades/:id', { id: ParamValue<true> }, { id: ParamValue<false> }>,
    '/school/parents': RouteRecordInfo<'/school/parents', '/school/parents', Record<never, never>, Record<never, never>>,
    '/school/school_years': RouteRecordInfo<'/school/school_years', '/school/school_years', Record<never, never>, Record<never, never>>,
    '/school/school_years/[id]': RouteRecordInfo<'/school/school_years/[id]', '/school/school_years/:id', { id: ParamValue<true> }, { id: ParamValue<false> }>,
    '/school/school_years/new': RouteRecordInfo<'/school/school_years/new', '/school/school_years/new', Record<never, never>, Record<never, never>>,
    '/school/students': RouteRecordInfo<'/school/students', '/school/students', Record<never, never>, Record<never, never>>,
    '/school/students/[id]': RouteRecordInfo<'/school/students/[id]', '/school/students/:id', { id: ParamValue<true> }, { id: ParamValue<false> }>,
    '/school/students/new': RouteRecordInfo<'/school/students/new', '/school/students/new', Record<never, never>, Record<never, never>>,
    '/school/subjects': RouteRecordInfo<'/school/subjects', '/school/subjects', Record<never, never>, Record<never, never>>,
    '/school/subjects/[id]': RouteRecordInfo<'/school/subjects/[id]', '/school/subjects/:id', { id: ParamValue<true> }, { id: ParamValue<false> }>,
    '/school/subjects/new': RouteRecordInfo<'/school/subjects/new', '/school/subjects/new', Record<never, never>, Record<never, never>>,
    '/settings/profile': RouteRecordInfo<'/settings/profile', '/settings/profile', Record<never, never>, Record<never, never>>,
    '/ui/': RouteRecordInfo<'/ui/', '/ui', Record<never, never>, Record<never, never>>,
  }
}

declare module 'vue-router/auto' {
  import type { RouteNamedMap } from 'vue-router/auto/routes'

  export type RouterTyped = _RouterTyped<RouteNamedMap>

  /**
   * Type safe version of `RouteLocationNormalized` (the type of `to` and `from` in navigation guards).
   * Allows passing the name of the route to be passed as a generic.
   */
  export type RouteLocationNormalized<Name extends keyof RouteNamedMap = keyof RouteNamedMap> = RouteLocationNormalizedTypedList<RouteNamedMap>[Name]

  /**
   * Type safe version of `RouteLocationNormalizedLoaded` (the return type of `useRoute()`).
   * Allows passing the name of the route to be passed as a generic.
   */
  export type RouteLocationNormalizedLoaded<Name extends keyof RouteNamedMap = keyof RouteNamedMap> = RouteLocationNormalizedLoadedTypedList<RouteNamedMap>[Name]

  /**
   * Type safe version of `RouteLocationResolved` (the returned route of `router.resolve()`).
   * Allows passing the name of the route to be passed as a generic.
   */
  export type RouteLocationResolved<Name extends keyof RouteNamedMap = keyof RouteNamedMap> = RouteLocationResolvedTypedList<RouteNamedMap>[Name]

  /**
   * Type safe version of `RouteLocation` . Allows passing the name of the route to be passed as a generic.
   */
  export type RouteLocation<Name extends keyof RouteNamedMap = keyof RouteNamedMap> = RouteLocationTypedList<RouteNamedMap>[Name]

  /**
   * Type safe version of `RouteLocationRaw` . Allows passing the name of the route to be passed as a generic.
   */
  export type RouteLocationRaw<Name extends keyof RouteNamedMap = keyof RouteNamedMap> =
    | RouteLocationAsString<RouteNamedMap>
    | RouteLocationAsRelativeTypedList<RouteNamedMap>[Name]
    | RouteLocationAsPathTypedList<RouteNamedMap>[Name]

  /**
   * Generate a type safe params for a route location. Requires the name of the route to be passed as a generic.
   */
  export type RouteParams<Name extends keyof RouteNamedMap> = RouteNamedMap[Name]['params']
  /**
   * Generate a type safe raw params for a route location. Requires the name of the route to be passed as a generic.
   */
  export type RouteParamsRaw<Name extends keyof RouteNamedMap> = RouteNamedMap[Name]['paramsRaw']

  export function useRouter(): RouterTyped
  export function useRoute<Name extends keyof RouteNamedMap = keyof RouteNamedMap>(name?: Name): RouteLocationNormalizedLoadedTypedList<RouteNamedMap>[Name]

  export const useLink: UseLinkFnTyped<RouteNamedMap>

  export function onBeforeRouteLeave(guard: NavigationGuard<RouteNamedMap>): void
  export function onBeforeRouteUpdate(guard: NavigationGuard<RouteNamedMap>): void

  export const RouterLink: RouterLinkTyped<RouteNamedMap>
  export const RouterLinkProps: RouterLinkPropsTyped<RouteNamedMap>

  // Experimental Data Fetching

  export function defineLoader<
    P extends Promise<any>,
    Name extends keyof RouteNamedMap = keyof RouteNamedMap,
    isLazy extends boolean = false,
  >(
    name: Name,
    loader: (route: RouteLocationNormalizedLoaded<Name>) => P,
    options?: _DefineLoaderOptions<isLazy>,
  ): _DataLoader<Awaited<P>, isLazy>
  export function defineLoader<
    P extends Promise<any>,
    isLazy extends boolean = false,
  >(
    loader: (route: RouteLocationNormalizedLoaded) => P,
    options?: _DefineLoaderOptions<isLazy>,
  ): _DataLoader<Awaited<P>, isLazy>

  export {
    _definePage as definePage,
    _HasDataLoaderMeta as HasDataLoaderMeta,
    _setupDataFetchingGuard as setupDataFetchingGuard,
    _stopDataFetchingScope as stopDataFetchingScope,
  } from 'unplugin-vue-router/runtime'
}

declare module 'vue-router' {
  import type { RouteNamedMap } from 'vue-router/auto/routes'

  export interface TypesConfig {
    beforeRouteUpdate: NavigationGuard<RouteNamedMap>
    beforeRouteLeave: NavigationGuard<RouteNamedMap>

    $route: RouteLocationNormalizedLoadedTypedList<RouteNamedMap>[keyof RouteNamedMap]
    $router: _RouterTyped<RouteNamedMap>

    RouterLink: RouterLinkTyped<RouteNamedMap>
  }
}
