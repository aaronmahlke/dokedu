/* eslint-disable */
import * as types from './graphql';
import type { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';

/**
 * Map of all GraphQL operations in the project.
 *
 * This map has several performance disadvantages:
 * 1. It is not tree-shakeable, so it will include all operations in the project.
 * 2. It is not minifiable, so the string of a GraphQL query will be multiple times inside the bundle.
 * 3. It does not support dead code elimination, so it will add unused operations.
 *
 * Therefore it is highly recommended to use the babel or swc plugin for production.
 */
const documents = {
    "\n    mutation generateFileURL($input: GenerateFileURLInput!) {\n      generateFileURL(input: $input) {\n        url\n      }\n    }\n  ": types.GenerateFileUrlDocument,
    "\n    mutation createFolder($input: CreateFolderInput!) {\n      createFolder(input: $input) {\n        id\n      }\n    }\n  ": types.CreateFolderDocument,
    "\n    query fileById($id: ID!) {\n      file(id: $id) {\n        id\n        name\n        parents {\n          id\n          name\n        }\n      }\n    }\n  ": types.FileByIdDocument,
    "\n    mutation singleUpload($input: FileUploadInput!) {\n      singleUpload(input: $input) {\n        id\n      }\n    }\n  ": types.SingleUploadDocument,
    "\n    query files($input: FilesFilterInput) {\n      files(input: $input) {\n        edges {\n          id\n          name\n          fileType\n          size\n        }\n      }\n    }\n  ": types.FilesDocument,
    "\n    query myFiles($input: MyFilesFilterInput) {\n      myFiles(input: $input) {\n        edges {\n          id\n          name\n          fileType\n          size\n        }\n      }\n    }\n  ": types.MyFilesDocument,
    "\n  mutation createEntry($input: CreateEntryInput!) {\n    createEntry(input: $input) {\n      id\n      date\n      body\n      deletedAt\n      user {\n        id\n        firstName\n        lastName\n      }\n      createdAt\n      tags {\n        id\n        name\n        color\n      }\n      events {\n        id\n        title\n      }\n      users {\n        id\n        firstName\n        lastName\n      }\n      userCompetences {\n        id\n        level\n        competence {\n          id\n          name\n          color\n          type\n        }\n      }\n    }\n  }\n": types.CreateEntryDocument,
    "\n  mutation updateEntry($input: UpdateEntryInput!) {\n    updateEntry(input: $input) {\n      id\n      date\n      body\n      deletedAt\n      user {\n        id\n        firstName\n        lastName\n      }\n      createdAt\n      tags {\n        id\n        name\n        color\n      }\n      events {\n        id\n        title\n      }\n      users {\n        id\n        firstName\n        lastName\n      }\n      userCompetences {\n        id\n        level\n        competence {\n          id\n          name\n          color\n          type\n        }\n      }\n    }\n  }\n": types.UpdateEntryDocument,
    "\n    query competences($search: String, $filter: CompetenceFilterInput) {\n      competences(search: $search, filter: $filter) {\n        edges {\n          id\n          name\n          type\n          color\n          grades\n          parents {\n            id\n            name\n            type\n            grades\n          }\n        }\n      }\n    }\n  ": types.CompetencesDocument,
    "\n    query events {\n      events {\n        edges {\n          id\n          title\n        }\n      }\n    }\n  ": types.EventsDocument,
    "\n    query users {\n      users(filter: { role: [student] }) {\n        edges {\n          id\n          firstName\n          lastName\n        }\n      }\n    }\n  ": types.UsersDocument,
    "\n    query entryById($id: ID!) {\n      entry(id: $id) {\n        id\n        date\n        body\n        deletedAt\n        user {\n          id\n          firstName\n          lastName\n        }\n        createdAt\n        tags {\n          id\n          name\n          color\n        }\n        events {\n          id\n          title\n        }\n        users {\n          id\n          firstName\n          lastName\n        }\n        userCompetences {\n          id\n          level\n          competence {\n            id\n            name\n            color\n            type\n            grades\n            parents {\n              id\n              name\n              grades\n            }\n          }\n        }\n      }\n    }\n  ": types.EntryByIdDocument,
    "\n    query event($id: ID!) {\n      event(id: $id) {\n        id\n        title\n        body\n        createdAt\n        startsAt\n        endsAt\n      }\n    }\n  ": types.EventDocument,
    "\n    query userById($id: ID!) {\n      user(id: $id) {\n        id\n        firstName\n        lastName\n        student {\n          id\n          grade\n          joinedAt\n          leftAt\n          entriesCount\n          competencesCount\n          eventsCount\n        }\n      }\n    }\n  ": types.UserByIdDocument,
    "\n    query subjectCompetences($subject: ID!, $user: ID!) {\n      competence(id: $subject) {\n        id\n        name\n        competences {\n          type\n          id\n          name\n          grades\n          parents {\n            id\n            name\n          }\n          competences {\n            id\n          }\n          userCompetences(userId: $user) {\n            id\n            level\n            entry {\n              id\n            }\n            createdBy {\n              firstName\n              lastName\n            }\n            createdAt\n          }\n        }\n      }\n    }\n  ": types.SubjectCompetencesDocument,
    "\n    mutation createUserCompetence($input: CreateUserCompetenceInput!) {\n      createUserCompetence(input: $input) {\n        id\n        level\n      }\n    }\n  ": types.CreateUserCompetenceDocument,
    "\n    query studentCompetences {\n      competences(filter: { type: subject }) {\n        edges {\n          id\n          name\n          grades\n        }\n      }\n    }\n  ": types.StudentCompetencesDocument,
    "\n    mutation CreateTag($input: CreateTagInput!) {\n      createTag(input: $input) {\n        id\n        name\n        color\n        deletedAt\n        createdAt\n      }\n    }\n  ": types.CreateTagDocument,
    "\n    mutation UpdateTag($id: ID!, $input: CreateTagInput!) {\n      updateTag(id: $id, input: $input) {\n        id\n        name\n        color\n        deletedAt\n        createdAt\n      }\n    }\n  ": types.UpdateTagDocument,
    "\n    mutation ArchiveTag($id: ID!) {\n      archiveTag(id: $id) {\n        id\n        name\n        color\n        deletedAt\n        createdAt\n      }\n    }\n  ": types.ArchiveTagDocument,
    "\n    query GetTags {\n      tags {\n        id\n        name\n        color\n        deletedAt\n        createdAt\n      }\n    }\n  ": types.GetTagsDocument,
    "\n  mutation archiveEntry($id: ID!) {\n    archiveEntry(id: $id) {\n      id\n    }\n  }\n": types.ArchiveEntryDocument,
    "\n    mutation signIn($email: String!, $password: String!) {\n        signIn(input: {email: $email, password: $password }) {\n            token\n        }\n    }\n": types.SignInDocument,
};

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 *
 *
 * @example
 * ```ts
 * const query = graphql(`query GetUser($id: ID!) { user(id: $id) { name } }`);
 * ```
 *
 * The query argument is unknown!
 * Please regenerate the types.
 */
export function graphql(source: string): unknown;

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation generateFileURL($input: GenerateFileURLInput!) {\n      generateFileURL(input: $input) {\n        url\n      }\n    }\n  "): (typeof documents)["\n    mutation generateFileURL($input: GenerateFileURLInput!) {\n      generateFileURL(input: $input) {\n        url\n      }\n    }\n  "];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation createFolder($input: CreateFolderInput!) {\n      createFolder(input: $input) {\n        id\n      }\n    }\n  "): (typeof documents)["\n    mutation createFolder($input: CreateFolderInput!) {\n      createFolder(input: $input) {\n        id\n      }\n    }\n  "];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query fileById($id: ID!) {\n      file(id: $id) {\n        id\n        name\n        parents {\n          id\n          name\n        }\n      }\n    }\n  "): (typeof documents)["\n    query fileById($id: ID!) {\n      file(id: $id) {\n        id\n        name\n        parents {\n          id\n          name\n        }\n      }\n    }\n  "];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation singleUpload($input: FileUploadInput!) {\n      singleUpload(input: $input) {\n        id\n      }\n    }\n  "): (typeof documents)["\n    mutation singleUpload($input: FileUploadInput!) {\n      singleUpload(input: $input) {\n        id\n      }\n    }\n  "];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query files($input: FilesFilterInput) {\n      files(input: $input) {\n        edges {\n          id\n          name\n          fileType\n          size\n        }\n      }\n    }\n  "): (typeof documents)["\n    query files($input: FilesFilterInput) {\n      files(input: $input) {\n        edges {\n          id\n          name\n          fileType\n          size\n        }\n      }\n    }\n  "];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query myFiles($input: MyFilesFilterInput) {\n      myFiles(input: $input) {\n        edges {\n          id\n          name\n          fileType\n          size\n        }\n      }\n    }\n  "): (typeof documents)["\n    query myFiles($input: MyFilesFilterInput) {\n      myFiles(input: $input) {\n        edges {\n          id\n          name\n          fileType\n          size\n        }\n      }\n    }\n  "];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation createEntry($input: CreateEntryInput!) {\n    createEntry(input: $input) {\n      id\n      date\n      body\n      deletedAt\n      user {\n        id\n        firstName\n        lastName\n      }\n      createdAt\n      tags {\n        id\n        name\n        color\n      }\n      events {\n        id\n        title\n      }\n      users {\n        id\n        firstName\n        lastName\n      }\n      userCompetences {\n        id\n        level\n        competence {\n          id\n          name\n          color\n          type\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  mutation createEntry($input: CreateEntryInput!) {\n    createEntry(input: $input) {\n      id\n      date\n      body\n      deletedAt\n      user {\n        id\n        firstName\n        lastName\n      }\n      createdAt\n      tags {\n        id\n        name\n        color\n      }\n      events {\n        id\n        title\n      }\n      users {\n        id\n        firstName\n        lastName\n      }\n      userCompetences {\n        id\n        level\n        competence {\n          id\n          name\n          color\n          type\n        }\n      }\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation updateEntry($input: UpdateEntryInput!) {\n    updateEntry(input: $input) {\n      id\n      date\n      body\n      deletedAt\n      user {\n        id\n        firstName\n        lastName\n      }\n      createdAt\n      tags {\n        id\n        name\n        color\n      }\n      events {\n        id\n        title\n      }\n      users {\n        id\n        firstName\n        lastName\n      }\n      userCompetences {\n        id\n        level\n        competence {\n          id\n          name\n          color\n          type\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  mutation updateEntry($input: UpdateEntryInput!) {\n    updateEntry(input: $input) {\n      id\n      date\n      body\n      deletedAt\n      user {\n        id\n        firstName\n        lastName\n      }\n      createdAt\n      tags {\n        id\n        name\n        color\n      }\n      events {\n        id\n        title\n      }\n      users {\n        id\n        firstName\n        lastName\n      }\n      userCompetences {\n        id\n        level\n        competence {\n          id\n          name\n          color\n          type\n        }\n      }\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query competences($search: String, $filter: CompetenceFilterInput) {\n      competences(search: $search, filter: $filter) {\n        edges {\n          id\n          name\n          type\n          color\n          grades\n          parents {\n            id\n            name\n            type\n            grades\n          }\n        }\n      }\n    }\n  "): (typeof documents)["\n    query competences($search: String, $filter: CompetenceFilterInput) {\n      competences(search: $search, filter: $filter) {\n        edges {\n          id\n          name\n          type\n          color\n          grades\n          parents {\n            id\n            name\n            type\n            grades\n          }\n        }\n      }\n    }\n  "];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query events {\n      events {\n        edges {\n          id\n          title\n        }\n      }\n    }\n  "): (typeof documents)["\n    query events {\n      events {\n        edges {\n          id\n          title\n        }\n      }\n    }\n  "];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query users {\n      users(filter: { role: [student] }) {\n        edges {\n          id\n          firstName\n          lastName\n        }\n      }\n    }\n  "): (typeof documents)["\n    query users {\n      users(filter: { role: [student] }) {\n        edges {\n          id\n          firstName\n          lastName\n        }\n      }\n    }\n  "];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query entryById($id: ID!) {\n      entry(id: $id) {\n        id\n        date\n        body\n        deletedAt\n        user {\n          id\n          firstName\n          lastName\n        }\n        createdAt\n        tags {\n          id\n          name\n          color\n        }\n        events {\n          id\n          title\n        }\n        users {\n          id\n          firstName\n          lastName\n        }\n        userCompetences {\n          id\n          level\n          competence {\n            id\n            name\n            color\n            type\n            grades\n            parents {\n              id\n              name\n              grades\n            }\n          }\n        }\n      }\n    }\n  "): (typeof documents)["\n    query entryById($id: ID!) {\n      entry(id: $id) {\n        id\n        date\n        body\n        deletedAt\n        user {\n          id\n          firstName\n          lastName\n        }\n        createdAt\n        tags {\n          id\n          name\n          color\n        }\n        events {\n          id\n          title\n        }\n        users {\n          id\n          firstName\n          lastName\n        }\n        userCompetences {\n          id\n          level\n          competence {\n            id\n            name\n            color\n            type\n            grades\n            parents {\n              id\n              name\n              grades\n            }\n          }\n        }\n      }\n    }\n  "];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query event($id: ID!) {\n      event(id: $id) {\n        id\n        title\n        body\n        createdAt\n        startsAt\n        endsAt\n      }\n    }\n  "): (typeof documents)["\n    query event($id: ID!) {\n      event(id: $id) {\n        id\n        title\n        body\n        createdAt\n        startsAt\n        endsAt\n      }\n    }\n  "];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query userById($id: ID!) {\n      user(id: $id) {\n        id\n        firstName\n        lastName\n        student {\n          id\n          grade\n          joinedAt\n          leftAt\n          entriesCount\n          competencesCount\n          eventsCount\n        }\n      }\n    }\n  "): (typeof documents)["\n    query userById($id: ID!) {\n      user(id: $id) {\n        id\n        firstName\n        lastName\n        student {\n          id\n          grade\n          joinedAt\n          leftAt\n          entriesCount\n          competencesCount\n          eventsCount\n        }\n      }\n    }\n  "];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query subjectCompetences($subject: ID!, $user: ID!) {\n      competence(id: $subject) {\n        id\n        name\n        competences {\n          type\n          id\n          name\n          grades\n          parents {\n            id\n            name\n          }\n          competences {\n            id\n          }\n          userCompetences(userId: $user) {\n            id\n            level\n            entry {\n              id\n            }\n            createdBy {\n              firstName\n              lastName\n            }\n            createdAt\n          }\n        }\n      }\n    }\n  "): (typeof documents)["\n    query subjectCompetences($subject: ID!, $user: ID!) {\n      competence(id: $subject) {\n        id\n        name\n        competences {\n          type\n          id\n          name\n          grades\n          parents {\n            id\n            name\n          }\n          competences {\n            id\n          }\n          userCompetences(userId: $user) {\n            id\n            level\n            entry {\n              id\n            }\n            createdBy {\n              firstName\n              lastName\n            }\n            createdAt\n          }\n        }\n      }\n    }\n  "];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation createUserCompetence($input: CreateUserCompetenceInput!) {\n      createUserCompetence(input: $input) {\n        id\n        level\n      }\n    }\n  "): (typeof documents)["\n    mutation createUserCompetence($input: CreateUserCompetenceInput!) {\n      createUserCompetence(input: $input) {\n        id\n        level\n      }\n    }\n  "];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query studentCompetences {\n      competences(filter: { type: subject }) {\n        edges {\n          id\n          name\n          grades\n        }\n      }\n    }\n  "): (typeof documents)["\n    query studentCompetences {\n      competences(filter: { type: subject }) {\n        edges {\n          id\n          name\n          grades\n        }\n      }\n    }\n  "];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation CreateTag($input: CreateTagInput!) {\n      createTag(input: $input) {\n        id\n        name\n        color\n        deletedAt\n        createdAt\n      }\n    }\n  "): (typeof documents)["\n    mutation CreateTag($input: CreateTagInput!) {\n      createTag(input: $input) {\n        id\n        name\n        color\n        deletedAt\n        createdAt\n      }\n    }\n  "];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation UpdateTag($id: ID!, $input: CreateTagInput!) {\n      updateTag(id: $id, input: $input) {\n        id\n        name\n        color\n        deletedAt\n        createdAt\n      }\n    }\n  "): (typeof documents)["\n    mutation UpdateTag($id: ID!, $input: CreateTagInput!) {\n      updateTag(id: $id, input: $input) {\n        id\n        name\n        color\n        deletedAt\n        createdAt\n      }\n    }\n  "];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation ArchiveTag($id: ID!) {\n      archiveTag(id: $id) {\n        id\n        name\n        color\n        deletedAt\n        createdAt\n      }\n    }\n  "): (typeof documents)["\n    mutation ArchiveTag($id: ID!) {\n      archiveTag(id: $id) {\n        id\n        name\n        color\n        deletedAt\n        createdAt\n      }\n    }\n  "];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query GetTags {\n      tags {\n        id\n        name\n        color\n        deletedAt\n        createdAt\n      }\n    }\n  "): (typeof documents)["\n    query GetTags {\n      tags {\n        id\n        name\n        color\n        deletedAt\n        createdAt\n      }\n    }\n  "];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation archiveEntry($id: ID!) {\n    archiveEntry(id: $id) {\n      id\n    }\n  }\n"): (typeof documents)["\n  mutation archiveEntry($id: ID!) {\n    archiveEntry(id: $id) {\n      id\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation signIn($email: String!, $password: String!) {\n        signIn(input: {email: $email, password: $password }) {\n            token\n        }\n    }\n"): (typeof documents)["\n    mutation signIn($email: String!, $password: String!) {\n        signIn(input: {email: $email, password: $password }) {\n            token\n        }\n    }\n"];

export function graphql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;