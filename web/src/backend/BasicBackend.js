// Copyright 2020 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import * as Setting from "../Setting";

export function getCommunityHealth() {
  return fetch(`${Setting.ServerUrl}/api/get-community-health`, {
    method: 'GET',
    credentials: 'include'
  }).then(res => res.json());
}

export function getForumVersion() {
  return fetch(`${Setting.ServerUrl}/api/get-forum-version`, {
    method: 'GET',
    credentials: 'include'
  }).then(res => res.json());
}

export function getOnlineNum() {
  return fetch(`${Setting.ServerUrl}/api/get-online-num`, {
    method: 'GET',
    credentials: 'include'
  }).then(res => res.json());
}

export function getNodeNavigation() {
  return fetch(`${Setting.ServerUrl}/api/node-navigation`, {
    method: 'GET',
    credentials: 'include'
  }).then(res => res.json());
}

export function getCaptcha() {
  return fetch(`${Setting.ServerUrl}/api/get-captcha`, {
    method: 'GET',
    credentials: 'include'
  }).then(res => res.json());
}

export function getValidateCode(phoneNumber) {
  return fetch(`${Setting.ServerUrl}/api/get-validate-code?phoneNumber=${phoneNumber}`, {
    method: 'GET',
    credentials: 'include'
  }).then(res => res.json());
}
