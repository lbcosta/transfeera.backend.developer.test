package queries

const Filter = `
SELECT q.id, q.status, q.name, q.document_number, q.email, q.pix_key_type, q.pix_key_value, q.bank, q.agency, q.account
FROM
    (SELECT lower(concat(b.name, b.status, b.pix_key_type, b.pix_key_value)) AS finder, b.* FROM beneficiaries b WHERE b.deleted_at IS NULL) q
WHERE q.finder SIMILAR TO lower('%${SEARCH}%')
`
