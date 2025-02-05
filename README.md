# Orrery - mTLS Bot

With mTLS bot you can easily create & validate mutual TLS (mTLS) connections between your services.

## Root Server

The root server is the server that will act as the root Certificate Authority (CA) for the mTLS bot. This
server will be in charge of generating intermediate CA's, server certificates, and client certificates.

### Environment Variables

| Name                       | Description                                           | Default |
|----------------------------|-------------------------------------------------------|---------|
| MTLS_BOT_STORAGE_DIRECTORY | Directory used to store data related to the mTLS bot. |         |
