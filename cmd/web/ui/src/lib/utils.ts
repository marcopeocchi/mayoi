export const getHost = (path: string) => import.meta.env.PROD
  ? `${window.location.href}${path.startsWith('/') ? path.substring(1) : path}`
  : `http://localhost:6969${path}`