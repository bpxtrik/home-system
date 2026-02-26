export interface Response {
  detail: string
}

export interface Motion {
  ID: string
  Timestamp: string
}

export interface GetResponse {
  data: Motion[]
  total_count: number
  total_pages: number
  page: number
  limit: number
}
