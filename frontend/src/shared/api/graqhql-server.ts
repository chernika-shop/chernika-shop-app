import { ApolloClient, InMemoryCache, ApolloLink } from '@apollo/client'

export const createApolloClient = () =>
  new ApolloClient({
    cache: new InMemoryCache(),
    link: new ApolloLink(),
  })
